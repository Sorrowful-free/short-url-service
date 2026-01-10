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

func ExampleHandlers_RegisterMakeShortJSONHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.UrlService.SetHasURLs(true)

	requestBody := model.ShortURLRequest{
		OriginalURL: "https://example.com/very/long/url/path",
	}
	jsonBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	var response model.ShortURLResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", response.ShortURL)

	// Output:
	// Status: 201
	// Short URL: http://localhost:8080/abc123
}

func ExampleHandlers_RegisterMakeShortJSONHandler_conflict() {

	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.UrlService.SetConflictURL("https://example.com/existing/url")

	requestBody := model.ShortURLRequest{
		OriginalURL: "https://example.com/existing/url",
	}
	jsonBody, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	var response model.ShortURLResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", response.ShortURL)

	// Output:
	// Status: 409
	// Short URL: http://localhost:8080/abc123
}
