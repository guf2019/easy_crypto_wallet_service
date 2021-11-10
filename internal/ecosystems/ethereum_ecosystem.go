package ecosystems

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/guf2019/easy_crypto_wallet_service/configs"
	"github.com/guf2019/easy_crypto_wallet_service/internal/utils"
	"golang.org/x/time/rate"
	"math/big"
	"sync"
)

type EthereumEcosystem struct {
	mu          sync.Mutex
	rateLimiter *rate.Limiter
	client      *ethclient.Client
}

func NewEthereumEcosystem(limiter *rate.Limiter) (*EthereumEcosystem, error) {
	client, err := ethclient.Dial(configs.EthereumNode)
	if err != nil {
		return nil, err
	}
	return &EthereumEcosystem{
		rateLimiter: limiter,
		client:      client,
	}, nil
}

func (k *EthereumEcosystem) GetBalance(ctx context.Context, address string) (*Balance, error) {
	addr := common.HexToAddress(address)
	balance, err := k.client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return &Balance{}, err
	}
	bigFloatBalance := new(big.Float)
	bigFloatBalance.SetString(balance.String())
	etherBalance := utils.WeiToEther(bigFloatBalance)
	floatEtherBalance, _ := etherBalance.Float64()
	return &Balance{Value: floatEtherBalance}, nil
}