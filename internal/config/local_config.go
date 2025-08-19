package config

import "flag"

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

	flag.StringVar(&localConfig.ListenAddr, "a", "localhost:8080", "listen address")
	flag.StringVar(&localConfig.BaseURL, "b", "http://localhost:8080", "base URL")
}
