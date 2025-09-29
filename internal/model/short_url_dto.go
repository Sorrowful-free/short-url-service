package model

type ShortURLDto struct {
	ShortUID    string
	OriginalURL string
	IsDeleted   bool
}

func NewShortURLDto(shortUID string, originalURL string, isDeleted bool) ShortURLDto {
	return ShortURLDto{
		ShortUID:    shortUID,
		OriginalURL: originalURL,
		IsDeleted:   isDeleted,
	}
}
