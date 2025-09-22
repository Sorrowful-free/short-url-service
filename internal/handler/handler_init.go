package handler

import (
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

const (
	MakeShortPath          = "/"
	MakeShortJSONPath      = "/api/shorten"
	MakeShortBatchJSONPath = "/api/shorten/batch"
	MakeOriginalPath       = "/:id"
	OriginalPathParam      = "id"
	PingDBPath             = "/ping"
	GetUserURLsPath        = "/api/user/urls"
)

type Handlers struct {
	internalEcho       *echo.Echo
	internalBaseURL    string
	internalURLService service.ShortURLService
}

func NewHandlers(echo *echo.Echo, baseURL string, urlService service.ShortURLService) (*Handlers, error) {

	return &Handlers{
		internalEcho:       echo,
		internalBaseURL:    baseURL,
		internalURLService: urlService,
	}, nil
}

func (h *Handlers) RegisterHandlers() *Handlers {
	h.RegisterMakeShortHandler()
	h.RegisterMakeOriginalHandler()
	h.RegisterMakeShortJSONHandler()
	h.RegisterMakeShortBatchJSONHandler()
	h.RegisterGetUserUrlsHandler()
	h.RegisterPingDBHandler()
	return h
}
