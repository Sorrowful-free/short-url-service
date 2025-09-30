package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPingDBHandler(t *testing.T) {
	t.Run("positive case ping database", func(t *testing.T) {
		e := echo.New()
		ctrl := gomock.NewController(t)
		urlService := mocks.NewMockShortURLService(ctrl)

		handlers, err := NewHandlers(e, consts.TestBaseURL, urlService)
		if err != nil {
			t.Fatalf("failed to create handlers: %v", err)
		}
		handlers.RegisterHandlers()

		urlService.EXPECT().Ping(gomock.Any()).Return(nil)

		req := httptest.NewRequest(http.MethodGet, PingDBPath, nil)
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code %d, received %d", http.StatusOK, resp.StatusCode)
	})

	t.Run("negative case ping database", func(t *testing.T) {
		e := echo.New()
		ctrl := gomock.NewController(t)
		urlService := mocks.NewMockShortURLService(ctrl)

		handlers, err := NewHandlers(e, consts.TestBaseURL, urlService)
		if err != nil {
			t.Fatalf("failed to create handlers: %v", err)
		}
		handlers.RegisterHandlers()

		urlService.EXPECT().Ping(gomock.Any()).Return(errors.New("test error"))

		req := httptest.NewRequest(http.MethodGet, PingDBPath, nil)
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "expected status code %d, received %d", http.StatusInternalServerError, resp.StatusCode)
	})

}
