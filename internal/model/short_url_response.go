package model

// ShortURLResponse represents the response payload containing a created short URL.
// It contains the full short URL that can be used to access the original URL.
type ShortURLResponse struct {
	ShortURL string `json:"result"` // The complete short URL (base URL + short identifier)
}
