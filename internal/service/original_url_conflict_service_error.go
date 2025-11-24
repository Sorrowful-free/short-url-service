package service

import "fmt"

// OriginalURLConflictServiceError represents an error that occurs when attempting
// to create a short URL for an original URL that already exists in the system.
type OriginalURLConflictServiceError struct {
	OriginalURL string // The original URL that caused the conflict
}

// NewOriginalURLConflictServiceError creates a new OriginalURLConflictServiceError.
// Parameters:
//   - originalURL: the original URL that already exists
//
// Returns an error instance.
func NewOriginalURLConflictServiceError(originalURL string) error {
	return &OriginalURLConflictServiceError{
		OriginalURL: originalURL,
	}
}

func (e *OriginalURLConflictServiceError) Error() string {
	return fmt.Sprintf("original url %s already exists", e.OriginalURL)
}
