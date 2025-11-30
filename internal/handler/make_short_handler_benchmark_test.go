package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func BenchmarkMakeShortHandler(b *testing.B) {
	e := echo.New()
	ctrl := gomock.NewController(b)
	urlService := mocks.NewMockShortURLService(ctrl)
	urlService.EXPECT().TryMakeShort(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), nil).
		AnyTimes()
	config := config.GetLocalConfig()
	handlers, err := NewHandlers(e, consts.TestBaseURL, urlService, config)
	if err != nil {
		b.Fatalf("failed to create handlers: %v", err)
	}
	handlers.RegisterHandlers()

	originalURL := consts.TestOriginalURL
	req := httptest.NewRequest(http.MethodPost, MakeShortPath, bytes.NewBufferString(originalURL))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		e.ServeHTTP(rr, req)
	}
}
