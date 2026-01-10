package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/golang/mock/gomock"
)

func BenchmarkMakeOriginalHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	urlService := testHandlers.urlService
	urlService.EXPECT().TryMakeOriginal(gomock.Any(), gomock.Any()).
		Return(model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), nil).
		AnyTimes()

	req := httptest.NewRequest(http.MethodGet, "/"+consts.TestShortUID, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
