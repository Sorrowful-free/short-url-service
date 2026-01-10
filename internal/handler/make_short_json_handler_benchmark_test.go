package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/golang/mock/gomock"
)

func BenchmarkMakeShortJSONHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	urlService := testHandlers.urlService
	urlService.EXPECT().TryMakeShort(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false), nil).
		AnyTimes()

	shortRequest := model.ShortURLRequest{
		OriginalURL: consts.TestOriginalURL,
	}
	jsonRequest, _ := json.Marshal(shortRequest)
	req := httptest.NewRequest(http.MethodPost, MakeShortJSONPath, bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
