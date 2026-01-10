package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
)

func ExampleHandlers_RegisterMakeOriginalHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo

	req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Location: %s\n", rec.Header().Get(consts.HeaderLocation))

	// Output:
	// Status: 307
	// Location: https://example.com/original-url
}

func ExampleHandlers_RegisterMakeOriginalHandler_deleted() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo

	req := httptest.NewRequest(http.MethodGet, "/deleted123", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 410
	// Body: short url is deleted
}
