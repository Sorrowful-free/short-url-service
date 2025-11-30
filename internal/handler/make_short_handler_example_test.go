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
	e := echo.New()

	urlService := &service.ExampleService{}

	config := config.GetLocalConfig()

	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

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
	urlService := &service.ExampleService{ConflictURL: "https://example.com/existing/url"}
	config := config.GetLocalConfig()

	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

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
