package config

import (
	"flag"
	"os"
)

type LocalConfig struct {
	ListenAddr string
	BaseURL    string
}

var localConfig *LocalConfig

func GetLocalConfig() *LocalConfig {
	return localConfig
}

func init() {
	localConfig = &LocalConfig{}

	//default values takes from flags
	flag.StringVar(&localConfig.ListenAddr, "a", "localhost:8080", "listen address")
	flag.StringVar(&localConfig.BaseURL, "b", "http://localhost:8080", "base URL")

	//override default values with values from environment variables if they are set
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress != "" {
		localConfig.ListenAddr = serverAddress
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL != "" {
		localConfig.BaseURL = baseURL
	}
}
