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
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMakeShortBatchJSONHandler(t *testing.T) {
	t.Run("positive case create short URL batch", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		urlService := testHandlers.urlService

		urlService.EXPECT().TryMakeShortBatch(gomock.Any(), gomock.Any(), gomock.Any()).Return([]model.ShortURLDto{model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), model.NewShortURLDto(consts.TestShortURL2, consts.TestOriginalURL2, false)}, nil)

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
		echo.ServeHTTP(rr, req)

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
