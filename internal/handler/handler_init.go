package handler

import "github.com/Sorrowful-free/short-url-service/internal/service"

var (
	internalUrlService service.ShortUrlService
)

func Init(urlService service.ShortUrlService) {
	internalUrlService = urlService
}
