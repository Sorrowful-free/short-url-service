package model

// BatchShortURLRequestDto represents a single item in a batch short URL creation request.
// It contains a correlation ID to match the request with the response, and the original URL.
type BatchShortURLRequestDto struct {
	CorrelationID string `json:"correlation_id"` // Unique identifier to correlate request and response
	OriginalURL   string `json:"original_url"`   // The original URL to be shortened
}

// BatchShortURLRequest represents a batch request for creating multiple short URLs.
// It is a slice of BatchShortURLRequestDto objects.
type BatchShortURLRequest []BatchShortURLRequestDto
