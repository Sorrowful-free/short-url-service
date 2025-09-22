package handler

import (
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

func TestGetUserUrlsHandler(t *testing.T) {
	t.Run("positive case get user URLs no content", func(t *testing.T) {
		e := echo.New()
		ctrl := gomock.NewController(t)
		urlService := mocks.NewMockShortURLService(ctrl)

		handlers, err := NewHandlers(e, urlService, consts.TestBaseURL, consts.TestUserIDCriptoKey)
		if err != nil {
			t.Fatalf("failed to create handlers: %v", err)
		}
		handlers.RegisterHandlers()

		urlService.EXPECT().GetUserUrls(gomock.Any(), gomock.Any()).Return([]model.ShortURLDto{}, nil)

		encryptedUserID, err := handlers.internalUserIDEncryptor.Encrypt(consts.TestUserID)
		if err != nil {
			t.Fatalf("failed to encrypt user ID: %v", err)
		}
		req := httptest.NewRequest(http.MethodGet, GetUserPath, nil)
		req.AddCookie(&http.Cookie{Name: UserIDCookieName, Value: encryptedUserID})
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNoContent, resp.StatusCode, "expected status code %d, received %d", http.StatusNoContent, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "no content", string(body), "expected body %s, received %s", "no content", string(body))
	})

	t.Run("positive case get user URLs any content", func(t *testing.T) {
		e := echo.New()
		ctrl := gomock.NewController(t)
		urlService := mocks.NewMockShortURLService(ctrl)

		handlers, err := NewHandlers(e, urlService, consts.TestBaseURL, consts.TestUserIDCriptoKey)
		if err != nil {
			t.Fatalf("failed to create handlers: %v", err)
		}
		handlers.RegisterHandlers()

		urlService.EXPECT().GetUserUrls(gomock.Any(), gomock.Any()).Return([]model.ShortURLDto{
			{
				ShortUID:    consts.TestShortUID,
				OriginalURL: consts.TestOriginalURL,
			},
		}, nil)

		encryptedUserID, err := handlers.internalUserIDEncryptor.Encrypt(consts.TestUserID)
		if err != nil {
			t.Fatalf("failed to encrypt user ID: %v", err)
		}
		req := httptest.NewRequest(http.MethodGet, GetUserPath, nil)
		req.AddCookie(&http.Cookie{Name: UserIDCookieName, Value: encryptedUserID})
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)

		resp := rr.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code %d, received %d", http.StatusOK, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		shortURL := model.UserShortURLResponse{}
		err = json.Unmarshal(body, &shortURL)
		if err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}
		assert.Equal(t, consts.TestShortURL, shortURL[0].ShortURL, "expected short URL %s, received %s", consts.TestShortUID, shortURL[0].ShortURL)
		assert.Equal(t, consts.TestOriginalURL, shortURL[0].OriginalURL, "expected original URL %s, received %s", consts.TestOriginalURL, shortURL[0].OriginalURL)
	})

}
