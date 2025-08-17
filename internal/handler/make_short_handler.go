package handler

import (
	"fmt"
	"io"
	"net/http"
)

func RegisterMakeShortHandler(mux *http.ServeMux) {
	mux.HandleFunc("POST /", makeShortHandlerInternal)
}

func makeShortHandlerInternal(w http.ResponseWriter, r *http.Request) {

	originalURL, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shortUID, err := internalURLService.TryMakeShort(string(originalURL))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shortURL := fmt.Sprintf("%s/%s", internalBaseURL, shortUID)
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, shortURL)

	fmt.Printf("process request for original URL:%s, with result:%s\n", originalURL, shortURL)
}
