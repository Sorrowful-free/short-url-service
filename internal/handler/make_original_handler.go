package handler

import (
	"net/http"
)

func MakeOriginalHandler() http.HandlerFunc {
	return makeShortHandlerInternal
}

func makeOriginalHandlerInternal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "unsuported method type", http.StatusBadRequest)
		return
	}
	shortUrl := r.PathValue("id")
	originalUrl, err := internalUrlService.TryMakeOriginal(shortUrl)
	if err != nil {
		// todo
	}
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Location", originalUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
