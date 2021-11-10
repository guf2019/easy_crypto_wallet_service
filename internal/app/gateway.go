package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/guf2019/easy_crypto_wallet_service/internal/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/guf2019/easy_crypto_wallet_service/insecure"
	pbExample "github.com/guf2019/easy_crypto_wallet_service/proto"
	"github.com/guf2019/easy_crypto_wallet_service/third_party"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// getOpenAPIHandler serves an OpenAPI UI.
func getOpenAPIHandler() http.Handler {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		panic("couldn't add extension type: " + err.Error())
	}
	// Use subdirectory in embedded files

	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

// StartGateway runs the gRPC-Gateway, dialling the provided address.
func StartGateway(dialAddr string, gatewayAddr string, log logger.Logger) error {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwmux := runtime.NewServeMux()
	err = pbExample.RegisterEasyCryptoWalletServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}
	oa := getOpenAPIHandler()

	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwmux.ServeHTTP(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/metrics") {
				promhttp.Handler().ServeHTTP(w, r)
				return
			}
			oa.ServeHTTP(w, r)
		}),
	}

	// Empty parameters mean use the TLS Config specified with the server.
	if strings.ToLower(os.Getenv("SERVE_HTTP")) == "true" {
		log.Infof("Serving gRPC-Gateway and OpenAPI Documentation on http://%v", gatewayAddr)
		return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
	}

	// #nosec G402
	gwServer.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{insecure.Cert},
	}
	log.Infof("Serving gRPC-Gateway and OpenAPI Documentation on https://%v", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServeTLS("", ""))
}
