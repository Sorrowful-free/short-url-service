package handler

import "github.com/Sorrowful-free/short-url-service/internal/service"

var (
	internalURLService service.ShortURLService
)

func Init(urlService service.ShortURLService) {
	internalURLService = urlService
}
