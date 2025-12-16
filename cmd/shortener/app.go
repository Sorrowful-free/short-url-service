package main

import (
	"net/http"
	_ "net/http/pprof"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"

	"context"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/middlewares"
	"github.com/Sorrowful-free/short-url-service/internal/repository"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

type App struct {
	internalContext         context.Context
	internalConfig          *config.LocalConfig
	internalLogger          logger.Logger
	internalUserIDEncryptor crypto.UserIDEncryptor
	internalEcho            *echo.Echo

	internalURLRepository repository.ShortURLRepository
	internalURLService    service.ShortURLService
}

func NewApp(ctx context.Context) *App {
	return &App{
		internalContext: ctx,
	}
}

func (a *App) InitLogger() error {
	l, err := logger.NewZapLogger()
	if err != nil {
		return err
	}
	a.internalLogger = l
	return nil
}

func (a *App) InitConfig() error {
	a.internalConfig = config.GetLocalConfig()
	return nil
}

func (a *App) InitUserIDEncryptor() error {
	userIDEncryptor, err := crypto.NewSha256UserIDEncryptor(a.internalConfig.UserIDCriptoKey)
	if err != nil {
		return err
	}
	a.internalUserIDEncryptor = userIDEncryptor
	return nil
}

func (a *App) InitURLRepository() error {

	var urlRepository repository.ShortURLRepository
	var err error

	if a.internalConfig.HasDatabaseDSN() {
		urlRepository, err = repository.NewPostgresShortURLRepository(a.internalConfig.DatabaseDSN, a.internalConfig.MigrationsPath, a.internalConfig.SkipMigrations)
	} else if a.internalConfig.HasFileStoragePath() {
		urlRepository, err = repository.NewFileStorageShortURLRepository(a.internalConfig.FileStoragePath)
	} else {
		urlRepository, err = repository.NewSimpleShortURLRepository(a.internalConfig.FileStoragePath)
	}

	if err != nil {
		return err
	}

	a.internalURLRepository = urlRepository
	return nil
}

func (a *App) InitURLService() error {
	urlService, err := service.NewSimpleService(a.internalConfig.UIDLength, a.internalConfig.UIDRetryCount, a.internalURLRepository, a.internalLogger)
	if err != nil {
		return err
	}
	a.internalURLService = urlService
	return nil
}

func (a *App) InitHandlers() error {
	e := echo.New()
	e.Use(middlewares.LoggerAsMiddleware(a.internalLogger))
	e.Use(middlewares.SimpleAuthMiddleware(a.internalUserIDEncryptor))
	e.Use(middlewares.GzipMiddleware(a.internalLogger))
	handlers, err := handler.NewHandlers(e, a.internalConfig.BaseURL, a.internalURLService, a.internalConfig)
	if err != nil {
		return err
	}
	handlers.RegisterHandlers()
	a.internalEcho = e
	return nil
}

func (a *App) Init() error {
	if err := a.InitLogger(); err != nil {
		return err
	}
	if err := a.InitConfig(); err != nil {
		return err
	}
	if err := a.InitUserIDEncryptor(); err != nil {
		return err
	}
	if err := a.InitURLRepository(); err != nil {
		return err
	}
	if err := a.InitURLService(); err != nil {
		return err
	}
	if err := a.InitHandlers(); err != nil {
		return err
	}
	return nil
}

func (a *App) PrintInfo(prefix string, str string) {
	if str == "" {
		str = "N/A"
	}
	a.internalLogger.Info(prefix + ": " + str)
}

func (a *App) Run() error {
	go func() {
		pprofAddr := "localhost:6060"
		a.internalLogger.Info("Starting pprof server on " + pprofAddr)
		if err := http.ListenAndServe(pprofAddr, nil); err != nil {
			a.internalLogger.Error("pprof server error: " + err.Error())
		}
	}()

	a.PrintInfo("Build version", buildVersion)
	a.PrintInfo("Build date", buildDate)
	a.PrintInfo("Build commit", buildCommit)

	if a.internalConfig.IsSecure {
		return a.internalEcho.StartTLS(a.internalConfig.ListenAddr, "cert.pem", "key.pem")
	} else {
		return a.internalEcho.Start(a.internalConfig.ListenAddr)
	}
}
