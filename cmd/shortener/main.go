package main

import (
	"flag"
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

	flag.Parse()

	localConfig := config.GetLocalConfig()

	urlService := service.NewFakeService()
	handler.Init(urlService, localConfig.BaseURL)

	e := echo.New()
	handler.RegisterMakeShortHandler(e)
	handler.RegisterMakeOriginalHandler(e)

	log.Printf("starting server and listening on addres %s ", localConfig.ListenAddr)
	return e.Start(localConfig.ListenAddr)

}
