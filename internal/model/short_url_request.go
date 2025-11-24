package model

// ShortURLRequest represents the request payload for creating a short URL.
// It contains the original URL to be shortened.
type ShortURLRequest struct {
	OriginalURL string `json:"url"` // The original URL to be shortened
}
