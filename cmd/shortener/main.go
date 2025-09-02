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
	l, err := logger.NewLogger()
	if err != nil {
		return err
	}
	e.Use(middlewares.LoggerAsMiddleware(l))
	e.Use(middlewares.GzipMiddleware)
	s := service.NewSimpleService(lc.UIDLength)
	handler.NewHandlers(e, s, lc.BaseURL).RegisterHandlers()
	return e.Start(lc.ListenAddr)
}
