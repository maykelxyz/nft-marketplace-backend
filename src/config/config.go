package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type ENVConfig struct {
	// App
	Port          string `env:"PORT"`
	BlockchainEnv string `env:"BLOCKCHAIN_ENV"`
	ReservoirKey  string `env:"RESERVOIR_API_KEY"`

	// Database
	DBHost     string `env:"DB_HOST"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`

	// Orderbook
	OrderbookSubmitURL          string `env:"ORDERBOOK_SUBMIT_URL"`
	OrderbookTLSSecureMode      bool   `env:"ORDERBOOK_TLS_SECURE_MODE"`
	OrderbookSSLCertificateFile string `env:"ORDERBOOK_SSL_CERTIFICATE_FILE_PATH"`
	OrderbookSSLKeyFile         string `env:"ORDERBOOK_SSL_KEY_FILE_PATH"`

	// CORS
	InOriginWhitelist  []string `env:"IN_ORIGIN_WHITELIST"`
	OutOriginWhitelist []string `env:"OUT_ORIGIN_WHITELIST"`
}

var envConfig *ENVConfig

func InitConfig(envFile string) (ENVConfig, error) {
	if envFile != "" {
		godotenv.Load(envFile)
	} else {
		godotenv.Load()
	}
	var cfg ENVConfig
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	// Define required variables
	requiredVars := map[string]string{
		"PORT":                                cfg.Port,
		"BLOCKCHAIN_ENV":                      cfg.BlockchainEnv,
		"RESERVOIR_API_KEY":                   cfg.ReservoirKey,
		"DB_HOST":                             cfg.DBHost,
		"DB_USER":                             cfg.DBUser,
		"DB_PASSWORD":                         cfg.DBPassword,
		"DB_NAME":                             cfg.DBName,
		"ORDERBOOK_SUBMIT_URL":                cfg.OrderbookSubmitURL,
		"ORDERBOOK_TLS_SECURE_MODE":           strconv.FormatBool(cfg.OrderbookTLSSecureMode),
		"ORDERBOOK_SSL_CERTIFICATE_FILE_PATH": cfg.OrderbookSSLCertificateFile,
		"ORDERBOOK_SSL_KEY_FILE_PATH":         cfg.OrderbookSSLKeyFile,
		"IN_ORIGIN_WHITELIST":                 strings.Join(cfg.InOriginWhitelist, ","),
		"OUT_ORIGIN_WHITELIST":                strings.Join(cfg.OutOriginWhitelist, ","),
	}

	// Check for missing required variables
	var missingVars []string
	for name, value := range requiredVars {
		if value == "" {
			missingVars = append(missingVars, name)
		}
	}

	if len(missingVars) > 0 {
		panic(fmt.Errorf("missing required environment variables: %s", strings.Join(missingVars, ", ")))
	}
	envConfig = &cfg
	return cfg, nil
}

func GetEnvConfig() *ENVConfig {
	return envConfig
}
