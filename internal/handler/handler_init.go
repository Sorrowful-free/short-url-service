package handler

import (
	"github.com/Sorrowful-free/short-url-service/internal/config"
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
	DeleteUserURLsPath     = "/api/user/urls"
)

// Handlers manages HTTP handlers for the short URL service.
// It encapsulates the Echo router, base URL, service layer, and configuration.
type Handlers struct {
	internalEcho       *echo.Echo
	internalBaseURL    string
	internalURLService service.ShortURLService
	internalConfig     *config.LocalConfig
}

// NewHandlers creates a new Handlers instance with the provided dependencies.
// Parameters:
//   - echo: the Echo router instance
//   - baseURL: the base URL for generating short URLs
//   - urlService: the service layer implementation
//   - config: the application configuration (currently unused)
//
// Returns a Handlers instance and an error if initialization fails.
func NewHandlers(echo *echo.Echo, baseURL string, urlService service.ShortURLService, config *config.LocalConfig) (*Handlers, error) {

	return &Handlers{
		internalEcho:       echo,
		internalBaseURL:    baseURL,
		internalURLService: urlService,
	}, nil
}

// RegisterHandlers registers all HTTP route handlers with the Echo router.
// This includes handlers for creating short URLs, retrieving original URLs,
// batch operations, user URL management, and database health checks.
// Returns the Handlers instance for method chaining.
func (h *Handlers) RegisterHandlers() *Handlers {
	h.RegisterMakeShortHandler()
	h.RegisterMakeOriginalHandler()
	h.RegisterMakeShortJSONHandler()
	h.RegisterMakeShortBatchJSONHandler()
	h.RegisterGetUserUrlsHandler()
	h.RegisterDeleteUserURLsHandler()
	h.RegisterPingDBHandler()
	return h
}
