package app

import (
	"context"
	"fmt"
	pbExample "github.com/guf2019/easy_crypto_wallet_service/proto"
	"google.golang.org/grpc/status"
)

// GetEthereumBalance This method helps get balance for address in ethereum blockchain.
func (b *Backend) CreateEthereumAccount(ctx context.Context, req *pbExample.CreateEthereumAccountRequest) (*pbExample.CreateEthereumAccountResponse, error) {
	password := req.GetPassword()
	request := fmt.Sprintf("CreateEthereumAccount request: password=%s.", password)
	keys, err := b.operationManager.CreateAccount(ctx, "ethereum", password)
	if err != nil {
		b.logger.Errorf("Error in CreateEthereumAccount function: %s", request)
		return nil, status.Errorf(500, "unknown error, ask administrator for help")
	}
	return &pbExample.CreateEthereumAccountResponse{PrivateKey: keys.PrivateKey, PublicKey: keys.PublicKey, Address: keys.Address}, nil
}
