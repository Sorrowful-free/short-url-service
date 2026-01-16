package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMakeOriginalHandler(t *testing.T) {
	t.Run("positive case make original URL", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		urlService := testHandlers.urlService

		urlService.EXPECT().TryMakeShort(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), nil)
		urlService.EXPECT().TryMakeOriginal(gomock.Any(), gomock.Any()).Return(model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), nil)
		originalURL := consts.TestOriginalURL
		req := httptest.NewRequest(http.MethodPost, MakeShortPath, bytes.NewBufferString(originalURL))
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected status code %d, received %d", http.StatusCreated, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		shortURL := string(body)

		assert.NotEmpty(t, shortURL, "short URL must be not empty")

		req = httptest.NewRequest(http.MethodGet, shortURL, nil)
		rr = httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp = rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode, "expected status code %d, received %d", http.StatusTemporaryRedirect, resp.StatusCode)
		assert.Equal(t, originalURL, resp.Header.Get("Location"), "expected location %s, received %s", originalURL, resp.Header.Get("Location"))
	})
	t.Run("positive case make original URL with is deleted", func(t *testing.T) {
		testHandlers := NewTestHandlers(t)
		echo := testHandlers.echo
		urlService := testHandlers.urlService

		urlService.EXPECT().TryMakeOriginal(gomock.Any(), gomock.Any()).Return(model.ShortURLDto{
			ShortUID:    consts.TestShortURL,
			OriginalURL: consts.TestOriginalURL,
			IsDeleted:   true,
		}, nil)

		req := httptest.NewRequest(http.MethodGet, MakeOriginalPath, nil)
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusGone, resp.StatusCode, "expected status code %d, received %d", http.StatusGone, resp.StatusCode)
	})
}
