package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/model"
)

func ExampleHandlers_RegisterMakeShortBatchJSONHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.URLService.SetHasURLs(true)

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
	echo.ServeHTTP(rec, req)

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
