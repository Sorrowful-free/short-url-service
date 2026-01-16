package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetStatHandler(t *testing.T) {
	t.Run("positive case get stats", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		statService := testHandlers.statService

		testHandlers.handlers.RegisterGetStatHandler()

		expectedStats := model.StatDto{
			Urls:  10,
			Users: 5,
		}

		statService.EXPECT().GetStats(gomock.Any()).Return(expectedStats, nil)

		req := httptest.NewRequest(http.MethodGet, GetUserURLsPath, nil)
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code %d, received %d", http.StatusOK, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		var stats model.StatDto
		err := json.Unmarshal(body, &stats)
		if err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}
		assert.Equal(t, expectedStats.Urls, stats.Urls, "expected urls %d, received %d", expectedStats.Urls, stats.Urls)
		assert.Equal(t, expectedStats.Users, stats.Users, "expected users %d, received %d", expectedStats.Users, stats.Users)
	})

	t.Run("negative case get stats with error", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		statService := testHandlers.statService

		testHandlers.handlers.RegisterGetStatHandler()

		expectedError := errors.New("database connection error")
		statService.EXPECT().GetStats(gomock.Any()).Return(model.StatDto{}, expectedError)

		req := httptest.NewRequest(http.MethodGet, GetUserURLsPath, nil)
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "expected status code %d, received %d", http.StatusInternalServerError, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, expectedError.Error(), string(body), "expected error message %s, received %s", expectedError.Error(), string(body))
	})
}
