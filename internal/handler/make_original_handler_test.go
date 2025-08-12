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

func TestMakeOriginalHandler(t *testing.T) {

	Init(service.NewFakeService())
	handler := MakeOriginalHandler()

	t.Run("POST method returns 400", func(t *testing.T) {
		internalTestMakeOriginalHandler(t, handler, http.MethodPost, "https://www.google.com", http.StatusBadRequest)
	})

	t.Run("GET method redirects", func(t *testing.T) {
		rr := internalTestMakeOriginalHandler(t, handler, http.MethodGet, "https://www.google.com", http.StatusTemporaryRedirect)
		assert.NotEmpty(t, rr.Header().Get("Location"), "Location header shouldn't be empty")
	})

}

func internalTestMakeOriginalHandler(t *testing.T, handler http.HandlerFunc, method string, body string, expectedCode int) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, "GET /{id}", io.NopCloser(strings.NewReader(body)))
	req.Header.Set("Content-Type", "text/plain")
	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != expectedCode {
		t.Errorf("expected status %d, got %d", expectedCode, rr.Code)
	}
	return rr
}
