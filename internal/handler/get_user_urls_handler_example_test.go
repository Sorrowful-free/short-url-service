package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func ExampleHandlers_RegisterGetUserUrlsHandler() {
	// Создаем экземпляр Echo
	e := echo.New()

	// Создаем мок сервиса (в реальном приложении используется реальный сервис)
	var urlService service.ShortURLService

	// Получаем конфигурацию
	config := config.GetLocalConfig()

	// Создаем хэндлеры
	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: GET запрос для получения всех URL пользователя
	req := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
	// В реальном приложении здесь должен быть установлен cookie с userID
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	var response model.UserShortURLResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Found %d URLs\n", len(response))
	for _, item := range response {
		fmt.Printf("ShortURL: %s, OriginalURL: %s\n", item.ShortURL, item.OriginalURL)
	}

	// Output:
	// Status: 200
	// Found 2 URLs
	// ShortURL: http://localhost:8080/abc123, OriginalURL: https://example.com/url1
	// ShortURL: http://localhost:8080/def456, OriginalURL: https://example.com/url2
}

func ExampleHandlers_RegisterGetUserUrlsHandler_noContent() {
	e := echo.New()
	var urlService service.ShortURLService
	config := config.GetLocalConfig()

	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: запрос для пользователя без сохраненных URL
	req := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", rec.Body.String())

	// Output:
	// Status: 204
	// Body: no content
}

