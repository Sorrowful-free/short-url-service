package repository

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

//go:generate mockgen --source=short_url_repository.go --destination=./../../mocks/mock_short_url_repository.go --package=mocks

// ShortURLRepository defines the interface for storing and retrieving short URLs.
// It provides methods for CRUD operations on short URL data.
type ShortURLRepository interface {
	// Save stores a new short URL associated with the given user ID.
	// Returns an error if the operation fails, including OriginalURLConflictRepositoryError
	// if the original URL already exists.
	Save(ctx context.Context, userID string, shortURL model.ShortURLDto) error

	// SaveBatch stores multiple short URLs in a single transaction.
	// Returns an error if the operation fails.
	SaveBatch(ctx context.Context, userID string, shortURLs []model.ShortURLDto) error

	// ContainsUID checks if a short URL with the given UID already exists.
	// Returns true if the UID exists, false otherwise.
	ContainsUID(ctx context.Context, shortUID string) bool

	// GetByUID retrieves a short URL by its unique identifier.
	// Returns the ShortURLDto and an error if not found.
	GetByUID(ctx context.Context, shortUID string) (model.ShortURLDto, error)

	// GetByOriginalURL retrieves a short URL by its original URL.
	// Returns the ShortURLDto and an error if not found.
	GetByOriginalURL(ctx context.Context, originalURL string) (model.ShortURLDto, error)

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
