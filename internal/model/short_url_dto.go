package model

type ShortURLDto struct {
	ShortUID    string
	OriginalURL string
}

func NewShortURLDto(shortUID string, originalURL string) ShortURLDto {
	return ShortURLDto{
		ShortUID:    shortUID,
		OriginalURL: originalURL,
	}
}
