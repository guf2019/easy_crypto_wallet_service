package ecosystems

import (
	"context"
	"errors"
)

type crypto_ecosystem interface {
	// GetBalance returns balance value from current crypto_ecosystem
	GetBalance(ctx context.Context, address string) (*Balance, error)
	// CreateAccount returns privatekey current crypto_ecosystem
	CreateAccount(ctx context.Context, password string) (*BlockchainKeys, error)
}

type EcosystemsInterface interface {
	GetBalance(ctx context.Context, ecosystemName, address string) (*Balance, error)
	CreateAccount(ctx context.Context, ecosystemName, password string) (*BlockchainKeys, error)
}

type Balance struct {
	Value float64
}

type BlockchainKeys struct {
	PrivateKey string
	PublicKey  string
	Address    string
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

func (ecosystems *CryptoEcosystems) CreateAccount(ctx context.Context, ecosystemName, password string) (*BlockchainKeys, error) {
	val, ok := ecosystems.ecosystems[ecosystemName]
	if !ok {
		return nil, errors.New("this ecosystem not supported")
	}
	keys, err := val.CreateAccount(ctx, password)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
