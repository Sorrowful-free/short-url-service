package main

import (
	"log"
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
)

func main() {

	err := run("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}

func run(address string) error {
	urlService := service.NewFakeService(address)
	handler.Init(urlService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handler.MakeShortHandler())
	mux.HandleFunc("GET /{id}", handler.MakeOriginalHandler())

	log.Printf("starting server and listening on addres %s ", address)
	err := http.ListenAndServe(address, mux)
	return err
}
