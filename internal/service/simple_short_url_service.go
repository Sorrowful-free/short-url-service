package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/repository"
)

// SimpleShortURLService is a concrete implementation of ShortURLService.
// It provides URL shortening functionality with configurable UID length and retry logic.
type SimpleShortURLService struct {
	uidLength          int
	retryCount         int
	logger             logger.Logger
	ShortURLRepository repository.ShortURLRepository
}

// NewSimpleService creates a new instance of SimpleShortURLService.
// Parameters:
//   - uidLength: the length of generated short URL identifiers
//   - retryCount: maximum number of retries when generating unique identifiers
//   - shortURLRepository: the repository implementation for storing and retrieving URLs
//   - logger: the logger instance for logging operations
//
// Returns a ShortURLService implementation and an error if initialization fails.
func NewSimpleService(uidLength int, retryCount int, shortURLRepository repository.ShortURLRepository, logger logger.Logger) (ShortURLService, error) {
	service := SimpleShortURLService{
		uidLength:          uidLength,
		retryCount:         retryCount,
		logger:             logger,
		ShortURLRepository: shortURLRepository,
	}
	return &service, nil
}

// TryMakeShort creates a new short URL for the given original URL.
// It generates a unique identifier and saves it to the repository.
// If the original URL already exists, it returns the existing short URL.
func (service SimpleShortURLService) TryMakeShort(ctx context.Context, userID string, originalURL string) (model.ShortURLDto, error) {

	shortUID, err := service.makeSimpleUUIDString(ctx)
	if err != nil {
		return model.ShortURLDto{}, fmt.Errorf("failed to make uid: %w", err)
	}

	dto := model.NewShortURLDto(shortUID, originalURL, false)

	service.logger.Info("short url created", "shortUID", shortUID, "originalURL", originalURL)

	err = service.ShortURLRepository.Save(ctx, userID, dto)
	if err != nil {
		var originalURLConflictError *repository.OriginalURLConflictRepositoryError
		if errors.As(err, &originalURLConflictError) {
			dto, err = service.ShortURLRepository.GetByOriginalURL(ctx, originalURLConflictError.OriginalURL)
			if err != nil {
				return model.ShortURLDto{}, fmt.Errorf("failed to get short url by original url: %w", err)
			}
			return dto, NewOriginalURLConflictServiceError(dto.OriginalURL)
		}

		return model.ShortURLDto{}, fmt.Errorf("failed to save short url: %w", err)
	}

	return dto, nil
}

// TryMakeOriginal retrieves the original URL for the given short URL identifier.
func (service SimpleShortURLService) TryMakeOriginal(ctx context.Context, shortUID string) (model.ShortURLDto, error) {
	dto, err := service.ShortURLRepository.GetByUID(ctx, shortUID)

	if err != nil {
		return model.ShortURLDto{}, fmt.Errorf("short url %s doesnot exist: %w", shortUID, err)
	}

	service.logger.Info("original url found", "shortUID", shortUID, "originalURL", dto.OriginalURL)

	return dto, nil
}

// TryMakeShortBatch creates multiple short URLs for the given list of original URLs.
func (service SimpleShortURLService) TryMakeShortBatch(ctx context.Context, userID string, originalURLs []string) ([]model.ShortURLDto, error) {
	shortURLs := make([]model.ShortURLDto, len(originalURLs))
	for i, originalURL := range originalURLs {
		shortUID, err := service.makeSimpleUUIDString(ctx)
		if err != nil {
			return nil, err
		}
		shortURLs[i] = model.NewShortURLDto(shortUID, originalURL, false)
	}

	err := service.ShortURLRepository.SaveBatch(ctx, userID, shortURLs)
	if err != nil {
		return nil, err
	}

	return shortURLs, nil
}

// GetUserUrls retrieves all short URLs associated with the given user ID.
func (service SimpleShortURLService) GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error) {
	return service.ShortURLRepository.GetUserUrls(ctx, userID)
}

// DeleteShortURLs marks the specified short URLs as deleted for the given user.
func (service SimpleShortURLService) DeleteShortURLs(ctx context.Context, userID string, shortURLs []string) error {
	return service.ShortURLRepository.DeleteShortURLs(ctx, userID, shortURLs)
}

// Ping checks the connectivity to the underlying data store.
func (service SimpleShortURLService) Ping(ctx context.Context) error {
	return service.ShortURLRepository.Ping(ctx)
}

func (service SimpleShortURLService) makeSimpleUUIDString(ctx context.Context) (string, error) {
	shortUID := ""
	err := error(nil)

	retryCount := service.retryCount
	for exist := true; exist; retryCount-- { //trying regenerate guid if it was allready registered

		if retryCount == 0 {
			return "", fmt.Errorf("failed to make uid: retry count exceeded")
		}

		shortUID, err = crypto.GenerateRandomSequenceString(service.uidLength)
		exist = service.ShortURLRepository.ContainsUID(ctx, shortUID)

		if err != nil {
			return "", fmt.Errorf("failed to make uid: %w", err)
		}
	}

	return shortUID, nil
}
