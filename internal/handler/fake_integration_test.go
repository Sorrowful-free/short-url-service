package handler

import (
	"io"
	"net/http"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestFakeIntegration(t *testing.T) {

	Init(service.NewFakeService())
	shortHandler := MakeShortHandler()
	originalHandler := MakeOriginalHandler()

	t.Run("check if all url will be pack and unpack", func(t *testing.T) {
		rr := internalTestMakeShortHandler(t, shortHandler, http.MethodPost, "https://www.google.com", http.StatusCreated)
		shortURL, _ := io.ReadAll(rr.Result().Body)
		rr.Result().Body.Close()
		assert.NotEmpty(t, shortURL, "short url must be exist")

		rr = internalTestMakeOriginalHandler(t, originalHandler, http.MethodGet, string(shortURL), http.StatusTemporaryRedirect)
		assert.NotEmpty(t, rr.Header().Get("Location"), "Location header shouldn't be empty")
		assert.NotNil(t, rr.Result().Body, "Location shouldn't be empty")
		rr.Result().Body.Close()

		originalUrl := rr.Result().Header.Get("Location")

		assert.Equal(t, "https://www.google.com", originalUrl, "url from location must be the same as original")
	})
}
