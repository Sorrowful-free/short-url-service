package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUserURLsHandler(t *testing.T) {
	t.Run("positive case delete user URLs", func(t *testing.T) {
		e := echo.New()
		ctrl := gomock.NewController(t)
		urlService := mocks.NewMockShortURLService(ctrl)
		config := config.GetLocalConfig()
		handlers, err := NewHandlers(e, consts.TestBaseURL, urlService, config)
		if err != nil {
			t.Fatalf("failed to create handlers: %v", err)
		}
		handlers.RegisterHandlers()

		urlService.EXPECT().DeleteShortURLs(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		deleteShortURLRequest := []string{consts.TestShortURL, consts.TestShortURL2}
		deleteShortURLRequestJSON, err := json.Marshal(deleteShortURLRequest)
		if err != nil {
			t.Fatalf("failed to marshal delete short URL request: %v", err)
		}

		req := httptest.NewRequest(http.MethodDelete, DeleteUserURLsPath, bytes.NewBuffer(deleteShortURLRequestJSON))
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusAccepted, resp.StatusCode, "expected status code %d, received %d", http.StatusAccepted, resp.StatusCode)
	})
}
