package handler

import (
	"fmt"
	"net/http"
)

func RegisterMakeOriginalHandler(mux *http.ServeMux) {
	mux.HandleFunc("GET /{id}", makeOriginalHandlerInternal)
}

func makeOriginalHandlerInternal(w http.ResponseWriter, r *http.Request) {
	shortUID := r.PathValue("id")
	originalURL, err := internalURLService.TryMakeOriginal(shortUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Header().Add("Location", originalURL)
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
	fmt.Printf("process request for short UID:%s, with result:%s\n", shortUID, originalURL)
}
