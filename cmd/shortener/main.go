package main

import (
	"log"
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
)

func main() {
	err := run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

func run(adress string) error {
	urlService := service.NewFakeService()
	handler.Init(urlService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handler.MakeShortHandler())
	mux.HandleFunc("GET /{id}", handler.MakeOriginalHandler())

	log.Printf("starting server and listening on addres %s ", adress)
	err := http.ListenAndServe(adress, mux)
	return err
}
