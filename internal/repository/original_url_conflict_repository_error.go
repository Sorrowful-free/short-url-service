package repository

import "fmt"

// OriginalURLConflictRepositoryError represents an error that occurs when attempting
// to save a short URL for an original URL that already exists in the repository.
// This typically happens when a unique constraint violation occurs in the database.
type OriginalURLConflictRepositoryError struct {
	OriginalURL string // The original URL that caused the conflict
}

// NewOriginalURLConflictRepositoryError creates a new OriginalURLConflictRepositoryError.
// Parameters:
//   - originalURL: the original URL that already exists
//
// Returns an error instance.
func NewOriginalURLConflictRepositoryError(originalURL string) error {
	return &OriginalURLConflictRepositoryError{
		OriginalURL: originalURL,
	}
}

func (e *OriginalURLConflictRepositoryError) Error() string {
	return fmt.Sprintf("original url %s already exists", e.OriginalURL)
}
