package service

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

//go:generate mockgen -source=short_url_service.go -destination=./../../mocks/mock_short_url_service.go --package=mocks

// ShortURLService defines the interface for URL shortening operations.
// It provides methods to create short URLs, retrieve original URLs,
// manage user URLs, and check database connectivity.
type ShortURLService interface {
	// TryMakeShort creates a new short URL for the given original URL.
	// If the original URL already exists, it returns the existing short URL
	// along with an OriginalURLConflictServiceError.
	// Returns the created or existing ShortURLDto and an error if the operation fails.
	TryMakeShort(ctx context.Context, userID string, originalURL string) (model.ShortURLDto, error)

	// TryMakeOriginal retrieves the original URL for the given short URL identifier.
	// Returns the ShortURLDto containing the original URL and an error if not found.
	TryMakeOriginal(ctx context.Context, shortURL string) (model.ShortURLDto, error)

	// TryMakeShortBatch creates multiple short URLs for the given list of original URLs.
	// Returns a slice of ShortURLDto objects and an error if the operation fails.
	TryMakeShortBatch(ctx context.Context, userID string, originalURLs []string) ([]model.ShortURLDto, error)

	// GetUserUrls retrieves all short URLs associated with the given user ID.
	// Returns a slice of ShortURLDto objects and an error if the operation fails.
	GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error)

	// DeleteShortURLs marks the specified short URLs as deleted for the given user.
	// Returns an error if the operation fails.
	DeleteShortURLs(ctx context.Context, userID string, shortURLs []string) error

	// Ping checks the connectivity to the underlying data store.
	// Returns an error if the connection cannot be established.
	Ping(ctx context.Context) error
}
