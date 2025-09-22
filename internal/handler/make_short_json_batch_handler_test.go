package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMakeShortBatchJSONHandler(t *testing.T) {
	t.Run("positive case create short URL batch", func(t *testing.T) {
		e := echo.New()
		ctrl := gomock.NewController(t)
		urlService := mocks.NewMockShortURLService(ctrl)

		handlers, err := NewHandlers(e, consts.TestBaseURL, urlService)
		if err != nil {
			t.Fatalf("failed to create handlers: %v", err)
		}
		handlers.RegisterHandlers()

		urlService.EXPECT().TryMakeShortBatch(gomock.Any(), gomock.Any(), gomock.Any()).Return([]string{consts.TestShortURL, consts.TestShortURL2}, nil)

		originalURL := consts.TestOriginalURL
		originalURL2 := consts.TestOriginalURL2
		shortRequest := model.BatchShortURLRequest{
			{
				CorrelationID: "1",
				OriginalURL:   originalURL,
			},
			{
				CorrelationID: "2",
				OriginalURL:   originalURL2,
			},
		}
		jsonRequest, _ := json.Marshal(shortRequest)
		req := httptest.NewRequest(http.MethodPost, MakeShortBatchJSONPath, bytes.NewBuffer(jsonRequest))
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, received %d", http.StatusCreated, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)

		shortResponse := model.BatchShortURLResponse{}
		json.Unmarshal(body, &shortResponse)
		shortURL := shortResponse[0].ShortURL
		shortURL2 := shortResponse[1].ShortURL
		correlationID := shortResponse[0].CorrelationID
		correlationID2 := shortResponse[1].CorrelationID

		assert.NotEmpty(t, shortURL, "short URL must be not empty")
		assert.NotEmpty(t, shortURL2, "short URL must be not empty")
		assert.Equal(t, correlationID, "1", "correlation ID must be 1")
		assert.Equal(t, correlationID2, "2", "correlation ID must be 2")
	})
}
