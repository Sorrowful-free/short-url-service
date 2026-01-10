package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
)

func ExampleHandlers_RegisterMakeShortHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo

	originalURL := "https://example.com/very/long/url/path"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(originalURL))
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 201
	// Short URL: http://localhost:8080/abc123
}

func ExampleHandlers_RegisterMakeShortHandler_conflict() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.UrlService.SetConflictUrl("https://example.com/existing/url")

	originalURL := "https://example.com/existing/url"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(originalURL))
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Short URL: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 409
	// Short URL: http://localhost:8080/abc123
}
