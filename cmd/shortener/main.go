package main

import (
	"log"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
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
	localConfig := config.GetLocalConfig()

	e := echo.New()
	handler.NewHandlers(e, service.NewSimpleService(localConfig.UIDLength), localConfig.BaseURL).RegisterHandlers()
	return e.Start(localConfig.ListenAddr)
}
