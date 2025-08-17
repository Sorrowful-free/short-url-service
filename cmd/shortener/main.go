package main

import (
	"log"
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
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

	mux := http.NewServeMux()
	handler.RegisterMakeShortHandler(mux)
	handler.RegisterMakeOriginalHandler(mux)

	log.Printf("starting server and listening on addres %s ", address)
	err := http.ListenAndServe(address, mux)
	return err
}
