package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPingDBHandler(t *testing.T) {
	t.Run("positive case ping database", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		urlService := testHandlers.urlService

		urlService.EXPECT().Ping(gomock.Any()).Return(nil)

		req := httptest.NewRequest(http.MethodGet, PingDBPath, nil)
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code %d, received %d", http.StatusOK, resp.StatusCode)
	})

	t.Run("negative case ping database", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		urlService := testHandlers.urlService

		urlService.EXPECT().Ping(gomock.Any()).Return(errors.New("test error"))

		req := httptest.NewRequest(http.MethodGet, PingDBPath, nil)
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "expected status code %d, received %d", http.StatusInternalServerError, resp.StatusCode)
	})

}
