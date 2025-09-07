package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMakeShortHandler(t *testing.T) {
	t.Run("positive case create short URL", func(t *testing.T) {
		e := echo.New()
		l, err := logger.NewLogger()
		if err != nil {
			t.Fatal(err)
		}
		NewHandlers(e, service.NewSimpleService(consts.TestUIDLength, consts.TestFileStoragePath, l), consts.TestBaseURL).RegisterHandlers()

		originalURL := consts.TestOriginalURL
		req := httptest.NewRequest(http.MethodPost, MakeShortPath, bytes.NewBufferString(originalURL))
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, received %d", http.StatusCreated, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		shortURL := string(body)

		assert.NotEmpty(t, shortURL, "short URL must be not empty")
	})
}
