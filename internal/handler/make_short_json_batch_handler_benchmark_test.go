package handler

import (
	"bytes"
	"encoding/json"
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

func BenchmarkMakeShortBatchJSONHandler(b *testing.B) {
	e := echo.New()
	ctrl := gomock.NewController(b)
	urlService := mocks.NewMockShortURLService(ctrl)
	urlService.EXPECT().TryMakeShortBatch(gomock.Any(), gomock.Any(), gomock.Any()).
		Return([]model.ShortURLDto{
			model.NewShortURLDto(consts.TestShortURL, consts.TestOriginalURL, false),
			model.NewShortURLDto(consts.TestShortURL2, consts.TestOriginalURL2, false),
		}, nil).
		AnyTimes()
	config := config.GetLocalConfig()
	handlers, err := NewHandlers(e, consts.TestBaseURL, urlService, config)
	if err != nil {
		b.Fatalf("failed to create handlers: %v", err)
	}
	handlers.RegisterHandlers()

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
		e.ServeHTTP(rr, req)
	}
}

