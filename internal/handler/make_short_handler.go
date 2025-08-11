package handler

import (
	"io"
	"net/http"
	"strconv"
)

func MakeShortHandler() http.HandlerFunc {
	return makeShortHandlerInternal
}

func makeShortHandlerInternal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "unsuported method type", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, "unsuported content type", http.StatusBadRequest)
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
	w.Header().Add("Content-Length", strconv.Itoa(len(shortUrl)))
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, shortUrl)

}
