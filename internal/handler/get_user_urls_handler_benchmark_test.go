package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func BenchmarkGetUserUrlsHandler(b *testing.B) {
	e := echo.New()
	ctrl := gomock.NewController(b)
	urlService := mocks.NewMockShortURLService(ctrl)
	urlService.EXPECT().GetUserUrls(gomock.Any(), gomock.Any()).
		Return([]model.ShortURLDto{
			{
				ShortUID:    consts.TestShortUID,
				OriginalURL: consts.TestOriginalURL,
			},
		}, nil).
		AnyTimes()

	userIDEncryptor, err := crypto.NewSha256UserIDEncryptor(consts.TestUserIDCriptoKey)
	if err != nil {
		b.Fatalf("failed to create user ID encryptor: %v", err)
	}

	config := config.GetLocalConfig()
	handlers, err := NewHandlers(e, consts.TestBaseURL, urlService, config)
	if err != nil {
		b.Fatalf("failed to create handlers: %v", err)
	}
	handlers.RegisterHandlers()

	encryptedUserID, err := userIDEncryptor.Encrypt(consts.TestUserID)
	if err != nil {
		b.Fatalf("failed to encrypt user ID: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, GetUserURLsPath, nil)
	req.AddCookie(&http.Cookie{Name: consts.UserIDCookieName, Value: encryptedUserID})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)
	}
}
