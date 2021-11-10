package app

import (
	"context"
	pbExample "github.com/guf2019/easy_crypto_wallet_service/proto"
	"google.golang.org/grpc/status"
)

// GetEthereumBalance This method helps get balance for address in ethereum blockchain.
func (b *Backend) GetEthereumBalance(ctx context.Context, req *pbExample.GetEthereumBalanceRequest) (*pbExample.GetEthereumBalanceResponse, error) {
	address := req.GetAddress()
	balance, err := b.operationManager.GetBalance(ctx, "ethereum", address)
	if err != nil {
		return nil, status.Errorf(500, "unknown error, ask administrator for help")
	}
	return &pbExample.GetEthereumBalanceResponse{Balance: balance, Address: address}, nil
}
