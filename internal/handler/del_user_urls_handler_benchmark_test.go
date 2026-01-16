package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/golang/mock/gomock"
)

func BenchmarkDeleteUserURLsHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	urlService := testHandlers.urlService
	urlService.EXPECT().DeleteShortURLs(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	deleteShortURLRequest := []string{consts.TestShortURL, consts.TestShortURL2}
	deleteShortURLRequestJSON, _ := json.Marshal(deleteShortURLRequest)
	req := httptest.NewRequest(http.MethodDelete, DeleteUserURLsPath, bytes.NewBuffer(deleteShortURLRequestJSON))
	req.Header.Set("Content-Type", "application/json")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
