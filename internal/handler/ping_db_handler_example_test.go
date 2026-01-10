package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
)

func ExampleHandlers_RegisterPingDBHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.URLService.SetPingError(false)

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 200
	// Body: pong
}

func ExampleHandlers_RegisterPingDBHandler_error() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.URLService.SetPingError(true)

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 500
	// Body: database connection error
}
