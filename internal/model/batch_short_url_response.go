package model

// BatchShortURLResponseDto represents a single item in a batch short URL creation response.
// It contains the correlation ID to match with the request and the created short URL.
type BatchShortURLResponseDto struct {
	CorrelationID string `json:"correlation_id"` // Unique identifier to correlate with the request
	ShortURL      string `json:"short_url"`      // The complete short URL that was created
}

// BatchShortURLResponse represents a batch response containing multiple created short URLs.
// It is a slice of BatchShortURLResponseDto objects.
type BatchShortURLResponse []BatchShortURLResponseDto
