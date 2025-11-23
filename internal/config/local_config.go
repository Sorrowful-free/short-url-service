package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

type LocalConfig struct {
	ListenAddr      string
	BaseURL         string
	UIDLength       int
	UIDRetryCount   int
	FileStoragePath string
	MigrationsPath  string
	SkipMigrations  bool
	DatabaseDSN     string
	UserIDLength    int
	UserIDCriptoKey string

	AuditFilePath string
	AuditURL      string
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
	flag.IntVar(&localConfig.UIDRetryCount, "r", 10, "retry count for the short URL")
	flag.StringVar(&localConfig.FileStoragePath, "f", "", "file storage path")
	flag.StringVar(&localConfig.MigrationsPath, "m", "file://./migrations", "migrations path")
	flag.BoolVar(&localConfig.SkipMigrations, "s", false, "skip migrations")
	flag.StringVar(&localConfig.DatabaseDSN, "d", "", "postgres DSN")
	flag.IntVar(&localConfig.UserIDLength, "u", 8, "length of the user ID")
	flag.StringVar(&localConfig.UserIDCriptoKey, "k", "", "user ID cripto key")
	flag.StringVar(&localConfig.AuditFilePath, "audit-file", "", "audit file path")
	flag.StringVar(&localConfig.AuditURL, "audit-url", "", "audit URL")
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

	fileStoragePath := os.Getenv("FILE_STORAGE_PATH")
	if fileStoragePath != "" {
		localConfig.FileStoragePath = fileStoragePath
	}

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath != "" {
		localConfig.MigrationsPath = migrationsPath
	}

	skipMigrations := os.Getenv("SKIP_MIGRATIONS")
	if skipMigrations != "" {
		localConfig.SkipMigrations = skipMigrations == "true"
	}

	databaseDSN := os.Getenv("DATABASE_DSN")
	if databaseDSN != "" {
		localConfig.DatabaseDSN = databaseDSN
	}

	userIDLength := os.Getenv("USER_ID_LENGTH")
	if userIDLength != "" {
		userIDLengthInt, err := strconv.Atoi(userIDLength)
		if err != nil {
			log.Fatalf("invalid USER_ID_LENGTH: %s", err)
		}
		localConfig.UserIDLength = userIDLengthInt
	}

	userIDCriptoKey := os.Getenv("USER_ID_CRIPTO_KEY")
	if userIDCriptoKey != "" {
		localConfig.UserIDCriptoKey = userIDCriptoKey
	}

	auditFilePath := os.Getenv("AUDIT_FILE")
	if auditFilePath != "" {
		localConfig.AuditFilePath = auditFilePath
	}

	auditURL := os.Getenv("AUDIT_URL")
	if auditURL != "" {
		localConfig.AuditURL = auditURL
	}

	return localConfig
}

func (c *LocalConfig) HasFileStoragePath() bool {
	return c.FileStoragePath != ""
}

func (c *LocalConfig) HasDatabaseDSN() bool {
	return c.DatabaseDSN != ""
}

func (c *LocalConfig) HasAuditFilePath() bool {
	if c == nil {
		return false
	}
	return c.AuditFilePath != ""
}

func (c *LocalConfig) HasAuditURL() bool {
	if c == nil {
		return false
	}
	return c.AuditURL != ""
}
