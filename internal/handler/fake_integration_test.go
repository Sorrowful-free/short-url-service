package handler

import (
	"io"
	"net/http"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestFakeIntegration(t *testing.T) {

	Init(service.NewFakeService("localhost:8080"))
	shortHandler := MakeShortHandler()
	originalHandler := MakeOriginalHandler()
	outsideShortURL := "empty"

	t.Run("check if all url will be pack", func(t *testing.T) {
		rr := internalTestMakeShortHandler(t, shortHandler, http.MethodPost, "https://www.google.com", http.StatusCreated)
		result := rr.Result()
		shortURL, _ := io.ReadAll(result.Body)

		assert.NotEmpty(t, shortURL, "short url must be exist")
		result.Body.Close()
		outsideShortURL = string(shortURL)
	})

	t.Run("check if original url will be unpack", func(t *testing.T) {

		rr := internalTestMakeOriginalHandler(t, originalHandler, http.MethodGet, outsideShortURL, http.StatusTemporaryRedirect)
		result := rr.Result()
		header := result.Header
		location := header.Get("Location")

		assert.NotEmpty(t, location, "Location header shouldn't be empty")
		assert.Equal(t, "https://www.google.com", location, "url from location must be the same as original")
		result.Body.Close()
	})
}
