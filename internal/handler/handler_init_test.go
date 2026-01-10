package handler

import (
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/Sorrowful-free/short-url-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

type TestHandlers struct {
	echo          *echo.Echo
	handlers      *Handlers
	config        *config.LocalConfig
	urlService    *mocks.MockShortURLService
	statService   *mocks.MockStatService
	urlRepository *mocks.MockShortURLRepository
}

type ExampleTestHandlers struct {
	Echo        *echo.Echo
	Handlers    *Handlers
	Config      *config.LocalConfig
	UrlService  *service.ExampleURLService
	StatService *service.ExampleStatService
}

func NewTestHandlers(t *testing.T) *TestHandlers {
	ctrl := gomock.NewController(t)
	echo := echo.New()
	config := config.GetLocalConfig()
	urlService := mocks.NewMockShortURLService(ctrl)
	statService := mocks.NewMockStatService(ctrl)
	urlRepository := mocks.NewMockShortURLRepository(ctrl)
	handlers, _ := NewHandlers(echo, urlService, statService, config)
	handlers.RegisterHandlers()
	return &TestHandlers{
		echo:          echo,
		handlers:      handlers,
		config:        config,
		urlService:    urlService,
		statService:   statService,
		urlRepository: urlRepository,
	}
}

func NewTestBenchmarkHandlers(b *testing.B) *TestHandlers {
	ctrl := gomock.NewController(b)
	echo := echo.New()
	config := config.GetLocalConfig()
	urlService := mocks.NewMockShortURLService(ctrl)
	statService := mocks.NewMockStatService(ctrl)
	urlRepository := mocks.NewMockShortURLRepository(ctrl)
	handlers, _ := NewHandlers(echo, urlService, statService, config)
	handlers.RegisterHandlers()
	return &TestHandlers{
		echo:          echo,
		handlers:      handlers,
		config:        config,
		urlService:    urlService,
		statService:   statService,
		urlRepository: urlRepository,
	}
}

func NewExampleHandlers() *ExampleTestHandlers {
	echo := echo.New()
	config := config.GetLocalConfig()
	urlService := &service.ExampleURLService{}
	statService := &service.ExampleStatService{}
	handlers, _ := NewHandlers(echo, urlService, statService, config)
	handlers.RegisterHandlers()
	return &ExampleTestHandlers{
		Echo:        echo,
		Handlers:    handlers,
		Config:      config,
		UrlService:  urlService,
		StatService: statService,
	}
}
