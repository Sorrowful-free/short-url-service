package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

type LocalConfig struct {
	ListenAddr string
	BaseURL    string
	UIDLength  int
}

var localConfig *LocalConfig

func GetLocalConfig() *LocalConfig {

	if localConfig != nil {
		return localConfig
	}

	localConfig = &LocalConfig{}

	//default values takes from flags
	flag.StringVar(&localConfig.ListenAddr, "a", "localhost:8080", "listen address")
	flag.StringVar(&localConfig.BaseURL, "b", "http://localhost:8080", "base URL")
	flag.IntVar(&localConfig.UIDLength, "l", 8, "length of the short URL")
	flag.Parse()

	//override default values with values from environment variables if they are set
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress != "" {
		localConfig.ListenAddr = serverAddress
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL != "" {
		localConfig.BaseURL = baseURL
	}

	uidLength := os.Getenv("UID_LENGTH")
	if uidLength != "" {
		uidLengthInt, err := strconv.Atoi(uidLength)
		if err != nil {
			log.Fatalf("invalid UID_LENGTH: %s", err)
		}
		localConfig.UIDLength = uidLengthInt
	}

	return localConfig
}
