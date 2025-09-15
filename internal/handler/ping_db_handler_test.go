package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPingDBHandler(t *testing.T) {
	t.Run("positive case ping database", func(t *testing.T) {
		e := echo.New()
		dbService := service.NewMockDBService(false)

		NewHandlers(e, nil, dbService, consts.TestBaseURL).RegisterHandlers()

		req := httptest.NewRequest(http.MethodGet, PingDBPath, nil)
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code %d, received %d", http.StatusOK, resp.StatusCode)
	})

	t.Run("negative case ping database", func(t *testing.T) {
		e := echo.New()
		dbService := service.NewMockDBService(true)

		NewHandlers(e, nil, dbService, consts.TestBaseURL).RegisterHandlers()

		req := httptest.NewRequest(http.MethodGet, PingDBPath, nil)
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "expected status code %d, received %d", http.StatusInternalServerError, resp.StatusCode)
	})

}
