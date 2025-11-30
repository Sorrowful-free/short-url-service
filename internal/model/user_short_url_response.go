package model

// UserShortURLResponseDto represents a single short URL in a user's URL list response.
// It contains both the short URL and the original URL for display purposes.
type UserShortURLResponseDto struct {
	ShortURL    string `json:"short_url"`    // The complete short URL
	OriginalURL string `json:"original_url"` // The original URL that was shortened
}

// UserShortURLResponse represents a response containing all short URLs for a user.
// It is a slice of UserShortURLResponseDto objects.
type UserShortURLResponse []UserShortURLResponseDto
