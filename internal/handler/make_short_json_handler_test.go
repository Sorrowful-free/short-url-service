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

func TestMakeShortJSONHandler(t *testing.T) {
	t.Run("positive case create short URL", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		urlService := testHandlers.urlService

		urlService.EXPECT().TryMakeShort(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), nil)

		originalURL := consts.TestOriginalURL
		shortRequest := model.ShortURLRequest{
			OriginalURL: originalURL,
		}
		jsonRequest, _ := json.Marshal(shortRequest)
		req := httptest.NewRequest(http.MethodPost, MakeShortJSONPath, bytes.NewBuffer(jsonRequest))
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, received %d", http.StatusCreated, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)

		shortResponse := model.ShortURLResponse{}
		json.Unmarshal(body, &shortResponse)
		shortURL := shortResponse.ShortURL

		assert.NotEmpty(t, shortURL, "short URL must be not empty")
	})
}
