package config

import (
	"os"
	"strconv"
)

type Config struct {
	APIKey    string
	APISecret string
	Testnet   bool
	Port      string
}

func Load() *Config {
	testnetStr := os.Getenv("BYBIT_TESTNET")
	testnet, _ := strconv.ParseBool(testnetStr)

	port := os.Getenv("MCP_PORT")
	if port == "" {
		port = "30001"
	}

	return &Config{
		APIKey:    os.Getenv("BYBIT_API_KEY"),
		APISecret: os.Getenv("BYBIT_API_SECRET"),
		Testnet:   testnet,
		Port:      port,
	}
}
