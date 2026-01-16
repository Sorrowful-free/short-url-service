package main

import (
	"net"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"

	"context"

	"github.com/Sorrowful-free/short-url-service/api"
	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/middlewares"
	"github.com/Sorrowful-free/short-url-service/internal/repository"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
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
	internalGRPCServer      *grpc.Server
	internalGRPCListener    net.Listener

	internalURLRepository repository.ShortURLRepository
	internalURLService    service.ShortURLService
	internalStatService   service.StatService
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

func (a *App) InitRepositories() error {

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

func (a *App) InitServices() error {
	urlService, err := service.NewSimpleService(a.internalConfig.UIDLength, a.internalConfig.UIDRetryCount, a.internalURLRepository, a.internalLogger)
	if err != nil {
		return err
	}
	a.internalURLService = urlService

	statService := service.NewStatService(a.internalURLRepository)
	a.internalStatService = statService

	return nil
}

func (a *App) InitHandlers() error {
	e := echo.New()
	e.Use(middlewares.LoggerAsMiddleware(a.internalLogger))
	e.Use(middlewares.SimpleAuthMiddleware(a.internalUserIDEncryptor))
	e.Use(middlewares.GzipMiddleware(a.internalLogger))
	handlers, err := handler.NewHandlers(e, a.internalURLService, a.internalStatService, a.internalConfig)
	if err != nil {
		return err
	}
	handlers.RegisterHandlers()
	a.internalEcho = e
	return nil
}

func (a *App) InitGRPCServer() error {
	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Create gRPC handler
	grpcHandler := handler.NewGRPCHandler(a.internalConfig.BaseURL, a.internalURLService, a.internalUserIDEncryptor)

	// Register service
	api.RegisterShortenerServiceServer(grpcServer, grpcHandler)

	a.internalGRPCServer = grpcServer

	// Create listener for gRPC server
	listener, err := net.Listen("tcp", a.internalConfig.GRPCListenAddr)
	if err != nil {
		return err
	}

	a.internalGRPCListener = listener
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
	if err := a.InitRepositories(); err != nil {
		return err
	}
	if err := a.InitServices(); err != nil {
		return err
	}
	if err := a.InitHandlers(); err != nil {
		return err
	}
	if err := a.InitGRPCServer(); err != nil {
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

	// Start gRPC server in a goroutine
	go func() {
		a.internalLogger.Info("Starting gRPC server on " + a.internalGRPCListener.Addr().String())
		if err := a.internalGRPCServer.Serve(a.internalGRPCListener); err != nil {
			a.internalLogger.Error("gRPC server error: " + err.Error())
		}
	}()

	a.PrintInfo("Build version", buildVersion)
	a.PrintInfo("Build date", buildDate)
	a.PrintInfo("Build commit", buildCommit)

	// Start HTTP server (blocking)
	if a.internalConfig.IsSecure {
		return a.internalEcho.StartTLS(a.internalConfig.ListenAddr, "cert.pem", "key.pem")
	} else {
		return a.internalEcho.Start(a.internalConfig.ListenAddr)
	}
}

func (a *App) Shutdown(ctx context.Context) error {
	a.internalLogger.Info("Начинаем штатное завершение сервера...")

	// Shutdown gRPC server
	if a.internalGRPCServer != nil {
		a.internalLogger.Info("Stopping gRPC server...")
		stopped := make(chan struct{})
		go func() {
			a.internalGRPCServer.GracefulStop()
			close(stopped)
		}()

		select {
		case <-stopped:
			a.internalLogger.Info("gRPC server stopped")
		case <-ctx.Done():
			a.internalLogger.Info("gRPC server force stop")
			a.internalGRPCServer.Stop()
		}
	}

	// Close gRPC listener
	if a.internalGRPCListener != nil {
		if err := a.internalGRPCListener.Close(); err != nil {
			a.internalLogger.Error("Error closing gRPC listener: " + err.Error())
		}
	}

	// Shutdown HTTP server
	return a.internalEcho.Shutdown(ctx)
}
