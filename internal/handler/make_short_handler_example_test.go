package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func ExampleHandlers_RegisterMakeShortHandler() {
	// Создаем экземпляр Echo
	e := echo.New()

	// Создаем мок сервиса (в реальном приложении используется реальный сервис)
	var urlService service.ShortURLService

	// Получаем конфигурацию
	config := config.GetLocalConfig()

	// Создаем хэндлеры
	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: POST запрос для создания короткого URL из оригинального URL
	originalURL := "https://example.com/very/long/url/path"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(originalURL))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 201
	// Short URL: http://localhost:8080/abc123
}

func ExampleHandlers_RegisterMakeShortHandler_conflict() {
	e := echo.New()
	var urlService service.ShortURLService
	config := config.GetLocalConfig()

	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: попытка создать короткий URL для уже существующего оригинального URL
	originalURL := "https://example.com/existing/url"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(originalURL))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 409
	// Short URL: http://localhost:8080/abc123
}

