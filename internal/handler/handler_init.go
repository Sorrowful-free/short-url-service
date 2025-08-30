package handler

import (
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

const (
	MakeShortPath         = "/"
	MakeOriginalPath      = "/:id"
	OriginalPathParam     = "id"
	HeaderContentType     = "Content-Type"
	HeaderContentTypeText = "text/plain; charset=utf-8"
	HeaderLocation        = "Location"
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
	return h
}
