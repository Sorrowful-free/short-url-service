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
	shortUrl := r.PathValue("id")
	originalUrl, err := internalUrlService.TryMakeOriginal(shortUrl)
	if err != nil {
		// todo
	}
	w.Header().Add("Content-Type", "text/plain")
	http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
}
