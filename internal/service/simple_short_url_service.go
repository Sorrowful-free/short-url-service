package service

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/repository"
)

type SimpleShortURLService struct {
	uidLength          int
	logger             logger.Logger
	ShortURLRepository repository.ShortURLRepository
}

func NewSimpleService(uidLength int, shortURLRepository repository.ShortURLRepository, logger logger.Logger) (ShortURLService, error) {
	service := SimpleShortURLService{
		uidLength:          uidLength,
		logger:             logger,
		ShortURLRepository: shortURLRepository,
	}
	return &service, nil
}

func (service SimpleShortURLService) TryMakeShort(ctx context.Context, originalURL string) (string, error) {

	shortUID, err := service.makeSimpleUUIDString(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to make uid: %w", err)
	}

	dto := model.New(shortUID, originalURL)

	service.logger.Info("short url created", "shortUID", shortUID, "originalURL", originalURL)

	err = service.ShortURLRepository.Save(ctx, dto)
	if err != nil {
		return "", fmt.Errorf("failed to save short url: %w", err)
	}

	return shortUID, nil
}

func (service SimpleShortURLService) TryMakeOriginal(ctx context.Context, shortUID string) (string, error) {
	dto, err := service.ShortURLRepository.GetByUID(ctx, shortUID)

	if err != nil {
		return "", fmt.Errorf("short url %s doesnot exist: %w", shortUID, err)
	}

	service.logger.Info("original url found", "shortUID", shortUID, "originalURL", dto.OriginalURL)

	return dto.OriginalURL, nil
}

func (service SimpleShortURLService) TryMakeShortBatch(ctx context.Context, originalURLs []string) ([]string, error) {
	shortURLs := make([]model.ShortURLDto, len(originalURLs))
	shortUIDs := make([]string, len(originalURLs))
	for i, originalURL := range originalURLs {
		shortUID, err := service.makeSimpleUUIDString(ctx)
		if err != nil {
			return nil, err
		}
		shortURLs[i] = model.New(shortUID, originalURL)
		shortUIDs[i] = shortUID
	}

	err := service.ShortURLRepository.SaveBatch(ctx, shortURLs)
	if err != nil {
		return nil, err
	}

	return shortUIDs, nil
}

func (service SimpleShortURLService) Ping(ctx context.Context) error {
	return service.ShortURLRepository.Ping(ctx)
}

func makeSimpleUIDString(uidLength int) (string, error) {
	b := make([]byte, uidLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}

func (service SimpleShortURLService) makeSimpleUUIDString(ctx context.Context) (string, error) {
	shortUID := ""
	err := error(nil)

	for exist := true; exist; { //trying regenerate guid if it wal allready registered
		shortUID, err = makeSimpleUIDString(service.uidLength)
		exist = service.ShortURLRepository.ContainsUID(ctx, shortUID)

		if err != nil {
			return "", fmt.Errorf("failed to make uid: %w", err)
		}
	}

	return shortUID, nil
}
