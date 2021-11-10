package operations

import (
	"github.com/guf2019/easy_crypto_wallet_service/internal/ecosystems"
	"github.com/guf2019/easy_crypto_wallet_service/internal/logger"
	"sync"
)

type Operations struct {
	mu         sync.Mutex
	logger     logger.Logger
	ecosystems ecosystems.EcosystemsInterface
}

func NewOperations(logger logger.Logger, ecosystems ecosystems.EcosystemsInterface) (*Operations, error) {
	return &Operations{
		logger:     logger,
		ecosystems: ecosystems,
	}, nil
}
