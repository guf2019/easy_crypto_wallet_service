package operations

import (
	"context"
	"github.com/guf2019/easy_crypto_wallet_service/internal/ecosystems"
)

func (o *Operations) GetBalance(ctx context.Context, ecosystem, address string) (float64, error) {
	balance, err := o.ecosystems.GetBalance(ctx, ecosystem, address)
	if err != nil {
		return 0, err
	}
	return balance.Value, nil
}

func (o *Operations) CreateAccount(ctx context.Context, ecosystem, password string) (*ecosystems.BlockchainKeys, error) {
	keys, err := o.ecosystems.CreateAccount(ctx, ecosystem, password)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
