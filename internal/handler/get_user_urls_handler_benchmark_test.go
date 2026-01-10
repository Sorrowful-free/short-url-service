package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/golang/mock/gomock"
)

func BenchmarkGetUserUrlsHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	urlService := testHandlers.urlService
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

	encryptedUserID, err := userIDEncryptor.Encrypt(consts.TestUserID)
	if err != nil {
		b.Fatalf("failed to encrypt user ID: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, GetUserURLsPath, nil)
	req.AddCookie(&http.Cookie{Name: consts.UserIDCookieName, Value: encryptedUserID})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
