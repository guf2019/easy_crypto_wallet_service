package ecosystems

import (
	"errors"
	"fmt"
	"github.com/guf2019/easy_crypto_wallet_service/configs"
	"github.com/guf2019/easy_crypto_wallet_service/internal/logger"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type CryptoEcosystems struct {
	mu         sync.Mutex
	log        logger.Logger
	ecosystems map[string]crypto_ecosystem
}

func New(log logger.Logger) (*CryptoEcosystems, error) {
	ecosystems := &CryptoEcosystems{
		log:        log,
		ecosystems: make(map[string]crypto_ecosystem),
	}

	err := ecosystems.initializeEcosystems()
	if err != nil {
		return nil,
			fmt.Errorf("unable to initialize crypto_ecosystem adapters: %s", err.Error())
	}

	return ecosystems, nil
}

func (ecosystems *CryptoEcosystems) initializeEcosystems() error {
	var err error
	ecosystems.log.Debugf("Starting CryptoEcosystems initialization")
	r := rate.Every(1 * time.Second)
	limiter := rate.NewLimiter(r, configs.EthereumEcosystemRateLimitInSecond)
	ecosystems.ecosystems["ethereum"], err = NewEthereumEcosystem(limiter)
	if err != nil {
		return err
	}
	ecosystems.log.Debugf("CryptoEcosystems initialized: %d", len(ecosystems.ecosystems))
	return nil
}

func (ecosystems *CryptoEcosystems) ecosystemByNameAndClient(ecosystemName string) (crypto_ecosystem, error) {
	ecosystem, ok := ecosystems.ecosystems[ecosystemName]
	if !ok {
		return nil, errors.New("Ecosystem not supported")
	}

	return ecosystem, nil
}
