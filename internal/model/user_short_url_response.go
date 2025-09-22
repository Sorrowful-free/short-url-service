package model

type UserShortURLResponseDto struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type UserShortURLResponse []UserShortURLResponseDto
