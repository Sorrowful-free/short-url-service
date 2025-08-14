package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestMakeShortHandler(t *testing.T) {

	Init(service.NewFakeService("localhost:8080"))
	handler := MakeShortHandler()

	t.Run("POST method returns 201", func(t *testing.T) {
		rr := internalTestMakeShortHandler(t, handler, http.MethodPost, "https://www.google.com", http.StatusCreated)
		result := rr.Result()
		assert.NotNil(t, result.Body, "Location header shouldn't be empty")
		result.Body.Close()
	})

	t.Run("GET method returns 400", func(t *testing.T) {
		internalTestMakeShortHandler(t, handler, http.MethodGet, "https://www.google.com", http.StatusBadRequest)

	})
}

func internalTestMakeShortHandler(t *testing.T, handler http.HandlerFunc, method string, body string, expectedCode int) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, "/", io.NopCloser(strings.NewReader(body)))
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != expectedCode {
		t.Errorf("expected status %d, got %d", expectedCode, rr.Code)
	}
	return rr
}
