package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/model"
)

func ExampleHandlers_RegisterDeleteUserURLsHandler() {
	handlers := handler.NewExampleHandlers()

	echo := handlers.Echo
	handlers.UrlService.SetHasURLs(true)

	deleteRequest := model.DeleteShortURLRequest{
		"abc123",
		"def456",
		"ghi789",
	}
	jsonBody, _ := json.Marshal(deleteRequest)

	req := httptest.NewRequest(http.MethodDelete, "/api/user/urls", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	echo.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 202
	// Body: accepted
}
