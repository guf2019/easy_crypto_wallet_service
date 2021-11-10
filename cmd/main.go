package main

import (
	"flag"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/guf2019/easy_crypto_wallet_service/insecure"
	"github.com/guf2019/easy_crypto_wallet_service/internal/app"
	"github.com/guf2019/easy_crypto_wallet_service/internal/config"
	"github.com/guf2019/easy_crypto_wallet_service/internal/ecosystems"
	"github.com/guf2019/easy_crypto_wallet_service/internal/logger"
	"github.com/guf2019/easy_crypto_wallet_service/internal/operations"
	pbExample "github.com/guf2019/easy_crypto_wallet_service/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	flag.Parse()

	// Set config and loggers
	conf := config.ReadConfig()
	log := logger.NewLogger(conf.LogFile, conf.LogLevel)

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard))

	addr := conf.GRPCApi.Host + ":" + strconv.Itoa(conf.GRPCApi.Port)
	grpcGatewayAddr := conf.RESTApi.Host + ":" + strconv.Itoa(conf.RESTApi.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf("Failed to listen:%v", err)
		os.Exit(1)
	}

	// Initialize exchanges
	ecosystems, err := ecosystems.New(log.WithModule("ecosystems"))
	if err != nil {
		log.Errorf("Unable to instantiate ecosystems: %s", err.Error())
		os.Exit(1)
	}

	// Initialize operation manager
	opers, err := operations.NewOperations(log.WithModule("operations"), ecosystems)
	if err != nil {
		log.Errorf("Unable to instantiate operations: %s", err.Error())
		os.Exit(1)
	}
	ourApp := app.New(log, opers)

	newServer := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	pbExample.RegisterEasyCryptoWalletServiceServer(newServer, ourApp)
	grpc_prometheus.Register(newServer)

	// Serve gRPC Server
	log.Infof("Serving gRPC on https://%v", addr)
	go func() {
		err = newServer.Serve(lis)
		if err != nil {
			log.Errorf("Failed to server gRPC:%v", err)
			os.Exit(1)
		}
	}()

	// Serve http gateway
	err = app.StartGateway("dns:///"+addr, grpcGatewayAddr, log)
	if err != nil {
		log.Errorf("Failed to start gateway:%v", err)
		os.Exit(1)
	}

	// Get termination signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
