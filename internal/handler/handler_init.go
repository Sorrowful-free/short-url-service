package handler

import "github.com/Sorrowful-free/short-url-service/internal/service"

var (
	internalURLService service.ShortURLService
	baseURL            string
)

func Init(urlService service.ShortURLService, baseURL string) {
	internalURLService = urlService
	baseURL = baseURL
}
