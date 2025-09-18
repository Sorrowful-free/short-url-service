package service

import "fmt"

type OriginalURLConflictServiceError struct {
	OriginalURL string
}

func NewOriginalURLConflictServiceError(originalURL string) error {
	return &OriginalURLConflictServiceError{
		OriginalURL: originalURL,
	}
}

func (e *OriginalURLConflictServiceError) Error() string {
	return fmt.Sprintf("original url %s already exists", e.OriginalURL)
}
