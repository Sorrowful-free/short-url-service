package handler

import (
	"net/http"
)

func MakeOriginalHandler() http.HandlerFunc {
	return makeOriginalHandlerInternal
}

func makeOriginalHandlerInternal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "unsuported method type", http.StatusBadRequest)
		return
	}
	shortURL := r.PathValue("id")
	originalURL, err := internalURLService.TryMakeOriginal(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Header().Add("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
