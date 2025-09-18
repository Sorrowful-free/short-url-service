package model

type BatchShortURLRequestDto struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchShortURLRequest []BatchShortURLRequestDto
