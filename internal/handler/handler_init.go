package handler

import (
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

const (
	MakeShortPath     = "/"
	MakeShortJSONPath = "/api/shorten"
	MakeOriginalPath  = "/:id"
	OriginalPathParam = "id"
)

type Handlers struct {
	internalEcho       *echo.Echo
	internalURLService service.ShortURLService
	internalBaseURL    string
}

func NewHandlers(echo *echo.Echo, urlService service.ShortURLService, baseURL string) *Handlers {
	return &Handlers{
		internalEcho:       echo,
		internalURLService: urlService,
		internalBaseURL:    baseURL,
	}
}

func (h *Handlers) RegisterHandlers() *Handlers {
	RegisterMakeShortHandler(h)
	RegisterMakeOriginalHandler(h)
	RegisterMakeShortJSONHandler(h)
	return h
}
