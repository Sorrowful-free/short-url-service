package main

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/middlewares"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/golang-migrate/migrate"
	"github.com/labstack/echo/v4"
)

type App struct {
	internalContext    context.Context
	internalConfig     *config.LocalConfig
	internalLogger     logger.Logger
	internalEcho       *echo.Echo
	internalURLService service.ShortURLService
	internalDBService  service.DBService
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

func (a *App) InitMigration() error {
	m, err := migrate.New("file:///migrations", a.internalConfig.DatabaseDSN)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) InitDBService() error {
	dbService, err := service.NewPostgresDBService(a.internalConfig.DatabaseDSN)
	if err != nil {
		return err
	}
	a.internalDBService = dbService
	return nil
}

func (a *App) InitURLService() error {
	urlService, err := service.NewSimpleService(a.internalConfig.UIDLength, a.internalConfig.FileStoragePath, a.internalLogger)
	if err != nil {
		return err
	}
	a.internalURLService = urlService
	return nil
}

func (a *App) InitHandlers() error {
	e := echo.New()
	e.Use(middlewares.LoggerAsMiddleware(a.internalLogger))
	e.Use(middlewares.GzipMiddleware(a.internalLogger))
	handlers := handler.NewHandlers(e, a.internalURLService, a.internalDBService, a.internalConfig.BaseURL)
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
	if err := a.InitMigration(); err != nil {
		return err
	}
	if err := a.InitDBService(); err != nil {
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

func (a *App) Run() error {
	return a.internalEcho.Start(a.internalConfig.ListenAddr)
}
