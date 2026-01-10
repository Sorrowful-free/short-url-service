package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/model"
)

func ExampleHandlers_RegisterGetStatHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.Handlers.RegisterGetStatHandler()
	handlers.StatService.SetGetStatsError(false)

	req := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	var stats model.StatDto
	body := rec.Body.String()
	json.Unmarshal([]byte(body), &stats)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("URLs: %d\n", stats.Urls)
	fmt.Printf("Users: %d\n", stats.Users)

	// Output:
	// Status: 200
	// URLs: 1
	// Users: 1
}

func ExampleHandlers_RegisterGetStatHandler_error() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.Handlers.RegisterGetStatHandler()
	handlers.StatService.SetGetStatsError(true)

	req := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil)
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 500
	// Body: database connection error
}
