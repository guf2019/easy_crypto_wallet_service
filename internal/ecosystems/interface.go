package ecosystems

import (
	"context"
	"errors"
)

type crypto_ecosystem interface {
	// GetBalance returns balance value current crypto_ecosystem
	GetBalance(ctx context.Context, address string) (*Balance, error)
}

type EcosystemsInterface interface {
	GetBalance(ctx context.Context, ecosystemName, address string) (*Balance, error)
}

type Balance struct {
	Value float64
}

func (ecosystems *CryptoEcosystems) GetBalance(ctx context.Context, ecosystemName, address string) (*Balance, error) {
	val, ok := ecosystems.ecosystems[ecosystemName]
	if !ok {
		return &Balance{}, errors.New("this ecosystem not supported")
	}
	balance, err := val.GetBalance(ctx, address)
	if err != nil {
		return &Balance{}, err
	}
	return balance, nil
}
