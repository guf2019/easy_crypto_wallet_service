package app

import (
	"context"
	"fmt"
	pbExample "github.com/guf2019/easy_crypto_wallet_service/proto"
	"google.golang.org/grpc/status"
)

// GetEthereumBalance This method helps get balance for address in ethereum blockchain.
func (b *Backend) GetEthereumBalance(ctx context.Context, req *pbExample.GetEthereumBalanceRequest) (*pbExample.GetEthereumBalanceResponse, error) {
	address := req.GetAddress()
	request := fmt.Sprintf("GetEthereumBalance request: address=%s.", address)
	balance, err := b.operationManager.GetBalance(ctx, "ethereum", address)
	if err != nil {
		b.logger.Errorf("Error in GetEthereumBalance function: %s", request)
		return nil, status.Errorf(500, "unknown error, ask administrator for help")
	}
	return &pbExample.GetEthereumBalanceResponse{Balance: balance, Address: address}, nil
}
