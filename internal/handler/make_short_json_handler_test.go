package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMakeShortJSONHandler(t *testing.T) {
	t.Run("positive case create short URL", func(t *testing.T) {
		e := echo.New()
		l, err := logger.NewLogger()
		if err != nil {
			t.Fatal(err)
		}
		NewHandlers(e, service.NewSimpleService(8, l), "http://localhost:8080").RegisterHandlers()

		originalURL := "http://example.com"
		shortRequest := model.ShortRequest{
			OriginalURL: originalURL,
		}
		jsonRequest, _ := json.Marshal(shortRequest)
		req := httptest.NewRequest(http.MethodPost, MakeShortJSONPath, bytes.NewBuffer(jsonRequest))
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, received %d", http.StatusCreated, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)

		shortResponse := model.ShortResponse{}
		json.Unmarshal(body, &shortResponse)
		shortURL := shortResponse.ShortURL

		assert.NotEmpty(t, shortURL, "short URL must be not empty")
	})
}
