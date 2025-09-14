package main

import (
	"log"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/middlewares"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	lc := config.GetLocalConfig()

	e := echo.New()
	l, err := logger.NewZapLogger()
	if err != nil {
		return err
	}
	e.Use(middlewares.LoggerAsMiddleware(l))
	e.Use(middlewares.GzipMiddleware(l))
	s, err := service.NewSimpleService(lc.UIDLength, lc.FileStoragePath, l)
	if err != nil {
		return err
	}

	dbService, err := service.NewPostgresDBService(lc.DatabaseDSN)
	if err != nil {
		return err
	}
	handler.NewHandlers(e, s, dbService, lc.BaseURL).RegisterHandlers()
	return e.Start(lc.ListenAddr)
}
