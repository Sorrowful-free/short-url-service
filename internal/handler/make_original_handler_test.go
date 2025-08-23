package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMakeOriginalHandler(t *testing.T) {
	t.Run("positive case make original URL", func(t *testing.T) {
		e := echo.New()
		NewHandlers(e, service.NewFakeService(8), "http://localhost:8080").RegisterHandlers()

		originalURL := "http://example.com"
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(originalURL))
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, received %d", http.StatusCreated, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		shortURL := string(body)

		assert.NotEmpty(t, shortURL, "short URL must be not empty")

		req = httptest.NewRequest(http.MethodGet, shortURL, nil)
		rr = httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp = rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode, "expected status code %d, received %d", http.StatusTemporaryRedirect, resp.StatusCode)
		assert.Equal(t, originalURL, resp.Header.Get("Location"), "expected location %s, received %s", originalURL, resp.Header.Get("Location"))
	})

}
