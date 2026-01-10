package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/model"
)

func ExampleHandlers_RegisterGetUserUrlsHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.URLService.SetHasURLs(true)

	req := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

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
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.URLService.SetHasURLs(false)

	req := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", rec.Body.String())

	// Output:
	// Status: 204
	// Body: no content
}
