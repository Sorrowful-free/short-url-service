package repository

import "fmt"

type OriginalURLConflictRepositoryError struct {
	OriginalURL string
}

func NewOriginalURLConflictRepositoryError(originalURL string) error {
	return &OriginalURLConflictRepositoryError{
		OriginalURL: originalURL,
	}
}

func (e *OriginalURLConflictRepositoryError) Error() string {
	return fmt.Sprintf("original url %s already exists", e.OriginalURL)
}
