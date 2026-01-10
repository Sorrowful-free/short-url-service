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

func BenchmarkMakeShortBatchJSONHandler(b *testing.B) {
	testHandlers := NewTestBenchmarkHandlers(b)
	echo := testHandlers.echo
	urlService := testHandlers.urlService
	urlService.EXPECT().TryMakeShortBatch(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]model.ShortURLDto{
			model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false),
			model.NewShortURLDto(consts.TestShortURL2, consts.TestOriginalURL2, false),
		}, nil).
		AnyTimes()

	shortRequest := model.BatchShortURLRequest{
		{
			CorrelationID: "1",
			OriginalURL:   consts.TestOriginalURL,
		},
		{
			CorrelationID: "2",
			OriginalURL:   consts.TestOriginalURL2,
		},
	}
	jsonRequest, _ := json.Marshal(shortRequest)
	req := httptest.NewRequest(http.MethodPost, MakeShortBatchJSONPath, bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		echo.ServeHTTP(rr, req)
	}
}
