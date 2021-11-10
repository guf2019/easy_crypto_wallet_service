package config

import (
	"fmt"
	"github.com/guf2019/easy_crypto_wallet_service/configs"
	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Paths to search files in, on order of priority, latest takes precedence
var configPaths = []string{"/etc/easy_crypto_wallet_service/", "/usr/local/etc/", "./configs/"}

// Config struct with field name
type Config struct {
	LogFile  string `yaml:"log_file" env:"EASY_CRYPTO_WALLET_SERVICE_LOG_FILE" env-default:"easy_crypto_wallet_service.log"`
	LogLevel string `yaml:"log_level" env:"EASY_CRYPTO_WALLET_SERVICE_LOG_LEVEL" env-default:"debug"`
	RESTApi  struct {
		Host     string   `yaml:"host" env:"EASY_CRYPTO_WALLET_SERVICE_REST_HOST" env-default:"0.0.0.0"`
		Port     int      `yaml:"port" env:"EASY_CRYPTO_WALLET_SERVICE_REST_PORT" env-default:"11000"`
		AuthKeys []string `yaml:"auth_keys" env:"EASY_CRYPTO_WALLET_SERVICE_REST_AUTH_KEYS" env-default:"notsecure"`
	} `yaml:"rest_api"`
	GRPCApi struct {
		Host string `yaml:"host" env:"EASY_CRYPTO_WALLET_SERVICE_REST_HOST" env-default:"0.0.0.0"`
		Port int    `yaml:"port" env:"EASY_CRYPTO_WALLET_SERVICE_REST_PORT" env-default:"10000"`
	} `yaml:"grpc_api"`
}

func ReadConfig() *Config {
	var conf Config

	for _, file := range getConfigFiles() {
		_ = cleanenv.ReadConfig(file, &conf)
	}

	_ = cleanenv.ReadEnv(&conf)

	return &conf
}

func getConfigFiles() []string {
	var filename string

	if filename == "" {
		filename = os.Getenv("EASY_CRYPTO_WALLET_SERVICE_CONFIG")
	}

	if filename != "" {
		return []string{filename}
	}

	files := make([]string, 0)
	for _, path := range configPaths {
		files = append(
			files,
			fmt.Sprintf("%s%s", path, configs.DefaultConfigFilename))
	}
	return files
}

func SaveConfig(config Config) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	for _, file := range getConfigFiles() {
		_, err := os.Stat(file)
		if err != nil {
			continue
		}
		err = ioutil.WriteFile(file, data, 0)
		if err != nil {
			return err
		}
	}
	return nil
}
