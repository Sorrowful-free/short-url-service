package handler_test

import (
	"bytes"
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

func ExampleHandlers_RegisterMakeShortJSONHandler() {
	// Создаем экземпляр Echo
	e := echo.New()

	// Создаем простую реализацию сервиса для примера
	urlService := &service.ExampleService{HasURLs: true}

	// Получаем конфигурацию
	config := config.GetLocalConfig()

	// Создаем хэндлеры
	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: POST запрос для создания короткого URL через JSON API
	requestBody := model.ShortURLRequest{
		OriginalURL: "https://example.com/very/long/url/path",
	}
	jsonBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	var response model.ShortURLResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", response.ShortURL)

	// Output:
	// Status: 201
	// Short URL: http://localhost:8080/abc123
}

func ExampleHandlers_RegisterMakeShortJSONHandler_conflict() {
	e := echo.New()
	urlService := &service.ExampleService{ConflictURL: "https://example.com/existing/url"}
	config := config.GetLocalConfig()

	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: попытка создать короткий URL для уже существующего оригинального URL
	requestBody := model.ShortURLRequest{
		OriginalURL: "https://example.com/existing/url",
	}
	jsonBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	var response model.ShortURLResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", response.ShortURL)

	// Output:
	// Status: 409
	// Short URL: http://localhost:8080/abc123
}
