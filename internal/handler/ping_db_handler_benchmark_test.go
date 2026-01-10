package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func BenchmarkPingDBHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	urlService := testHandlers.urlService
	urlService.EXPECT().Ping(gomock.Any()).
		Return(nil).
		AnyTimes()

	req := httptest.NewRequest(http.MethodGet, PingDBPath, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
