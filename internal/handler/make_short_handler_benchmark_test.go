package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/golang/mock/gomock"
)

func BenchmarkMakeShortHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	urlService := testHandlers.urlService
	urlService.EXPECT().TryMakeShort(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), nil).
		AnyTimes()

	originalURL := consts.TestOriginalURL
	req := httptest.NewRequest(http.MethodPost, MakeShortPath, bytes.NewBufferString(originalURL))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
