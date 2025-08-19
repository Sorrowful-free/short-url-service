package main

import (
	"log"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {

	err := run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func run(address string) error {
	urlService := service.NewFakeService()
	handler.Init(urlService, address)

	e := echo.New()
	handler.RegisterMakeShortHandler(e)
	handler.RegisterMakeOriginalHandler(e)

	log.Printf("starting server and listening on addres %s ", address)
	return e.Start(address)

}
