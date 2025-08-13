package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestMakeOriginalHandler(t *testing.T) {

	Init(service.NewFakeService())
	handler := MakeOriginalHandler()

	t.Run("POST method returns 400", func(t *testing.T) {
		internalTestMakeOriginalHandler(t, handler, http.MethodPost, "C754D531", http.StatusBadRequest)
	})

	t.Run("GET method redirects", func(t *testing.T) {
		rr := internalTestMakeOriginalHandler(t, handler, http.MethodGet, "C754D531", http.StatusTemporaryRedirect)
		assert.NotEmpty(t, rr.Result().Header.Get("Location"), "Location header shouldn't be empty")
		rr.Result().Body.Close()
	})

}

func internalTestMakeOriginalHandler(t *testing.T, handler http.HandlerFunc, method string, shortUrl string, expectedCode int) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, "/", nil)
	req.SetPathValue("id", shortUrl)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	rr := httptest.NewRecorder()
	handler(rr, req)

	assert.Equal(t, rr.Code, expectedCode, "expected status %d, got %d", expectedCode, rr.Code)

	return rr
}
