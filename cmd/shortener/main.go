package main

import (
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
)

func main() {
	urlService := service.NewFakeService()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handler.MakeShortHandler(urlService))
	mux.HandleFunc("GET /{id}", handler.MakeOriginalHandler(urlService))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
