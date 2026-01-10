package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/golang/mock/gomock"
)

func BenchmarkGetStatHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	statService := testHandlers.statService

	testHandlers.handlers.RegisterGetStatHandler()

	expectedStats := model.StatDto{
		Urls:  10,
		Users: 5,
	}

	statService.EXPECT().GetStats(gomock.Any()).
		Return(expectedStats, nil).
		AnyTimes()

	req := httptest.NewRequest(http.MethodGet, GetUserURLsPath, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
