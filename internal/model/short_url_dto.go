package model

type ShortURLDto struct {
	ShortURL    string
	OriginalURL string
}

func New(shortURL string, originalURL string) ShortURLDto {
	return ShortURLDto{
		ShortURL:    shortURL,
		OriginalURL: originalURL,
	}
}
