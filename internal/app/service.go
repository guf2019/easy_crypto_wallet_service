package app

import (
	"github.com/guf2019/easy_crypto_wallet_service/internal/logger"
	"github.com/guf2019/easy_crypto_wallet_service/internal/operations"
)

// Backend implements the protobuf interface
type Backend struct {
	logger           logger.Logger
	operationManager *operations.Operations
}

// New initializes a new Backend struct.
func New(log logger.Logger, opers *operations.Operations) *Backend {
	return &Backend{
		logger:           log,
		operationManager: opers,
	}
}
