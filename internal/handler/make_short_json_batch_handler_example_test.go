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

func ExampleHandlers_RegisterMakeShortBatchJSONHandler() {
	// Создаем экземпляр Echo
	e := echo.New()

	// Создаем мок сервиса (в реальном приложении используется реальный сервис)
	var urlService service.ShortURLService

	// Получаем конфигурацию
	config := config.GetLocalConfig()

	// Создаем хэндлеры
	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: POST запрос для создания нескольких коротких URL через batch API
	batchRequest := model.BatchShortURLRequest{
		{
			CorrelationID: "1",
			OriginalURL:   "https://example.com/url1",
		},
		{
			CorrelationID: "2",
			OriginalURL:   "https://example.com/url2",
		},
		{
			CorrelationID: "3",
			OriginalURL:   "https://example.com/url3",
		},
	}
	jsonBody, _ := json.Marshal(batchRequest)

	req := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	var response []model.BatchShortURLResponseDto
	json.Unmarshal(rec.Body.Bytes(), &response)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Created %d short URLs\n", len(response))
	for _, item := range response {
		fmt.Printf("CorrelationID: %s, ShortURL: %s\n", item.CorrelationID, item.ShortURL)
	}

	// Output:
	// Status: 201
	// Created 3 short URLs
	// CorrelationID: 1, ShortURL: http://localhost:8080/abc123
	// CorrelationID: 2, ShortURL: http://localhost:8080/def456
	// CorrelationID: 3, ShortURL: http://localhost:8080/ghi789
}

