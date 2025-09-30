package model

import "github.com/google/uuid"

type ShortURLSafeDto struct {
	UUID        string `json:"uuid"`
	ShortUID    string `json:"short_url"` // to be honest that naming not fine to me but it will be checked by autotests
	OriginalURL string `json:"original_url"`
	IsDeleted   bool   `json:"is_deleted"`
}

func NewShortURLSafeDto(ShortURLDto ShortURLDto) ShortURLSafeDto {
	return ShortURLSafeDto{
		UUID:        uuid.New().String(),
		ShortUID:    ShortURLDto.ShortUID,
		OriginalURL: ShortURLDto.OriginalURL,
		IsDeleted:   ShortURLDto.IsDeleted,
	}
}
