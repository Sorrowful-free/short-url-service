package repository

import "fmt"

type OriginalURLConflictRepositoryError struct {
	ShortURL    string
	OriginalURL string
}

func NewOriginalURLConflictRepositoryError(shortURL, originalURL string) error {
	return &OriginalURLConflictRepositoryError{
		ShortURL:    shortURL,
		OriginalURL: originalURL,
	}
}

func (e *OriginalURLConflictRepositoryError) Error() string {
	return fmt.Sprintf("original url %s already exists for short url %s", e.OriginalURL, e.ShortURL)
}
