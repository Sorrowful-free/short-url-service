package handler

import (
	"fmt"
	"io"
	"net/http"
)

func MakeShortHandler() http.HandlerFunc {
	return makeShortHandlerInternal
}

func makeShortHandlerInternal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "unsuported method type", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
		errorMessage := fmt.Sprintf("unsuported content type %s", r.Header.Get("Content-Type"))
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}
	originalUrl, err := io.ReadAll(r.Body)
	if err != nil {

	}
	shortUrl, err := internalUrlService.TryMakeShort(string(originalUrl))
	if err != nil {
		// todo
	}
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, shortUrl)
}
