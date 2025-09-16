package model

type BatchShortURLResponseDto struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

type BatchShortURLResponse []BatchShortURLResponseDto
