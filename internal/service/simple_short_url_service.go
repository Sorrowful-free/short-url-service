package service

import (
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

func NewSimpleService(uidLength int, fileStoragePath string, logger logger.Logger) (ShortURLService, error) {
	shortURLRepository, err := repository.NewSimpleShortURLRepository(fileStoragePath)
	if err != nil {
		return nil, err
	}
	service := SimpleShortURLService{
		uidLength:          uidLength,
		logger:             logger,
		ShortURLRepository: shortURLRepository,
	}
	return &service, nil
}

func (service SimpleShortURLService) TryMakeShort(originalURL string) (string, error) {

	shortUID := ""
	err := error(nil)

	for exist := true; exist; { //trying regenerate guid if it wal allready registered
		shortUID, err = makeSimpleUIDString(service.uidLength)
		_, exist = service.shortUIDs[shortUID]

		if err != nil {
			return "", fmt.Errorf("failed to make uid: %w", err)
		}
	}

	dto := model.New(shortUID, originalURL)

	service.shortUIDs[shortUID] = dto
	service.logger.Info("short url created", "shortUID", shortUID, "originalURL", originalURL)

	err = service.ShortURLRepository.Save(dto)
	if err != nil {
		return "", fmt.Errorf("failed to save short url: %w", err)
	}

	return shortUID, nil
}

func (service SimpleShortURLService) TryMakeOriginal(shortUID string) (string, error) {
	dto, exist := service.shortUIDs[shortUID]

	if !exist {
		return "", fmt.Errorf("short url %s doesnot exist ", shortUID)
	}

	service.logger.Info("original url found", "shortUID", shortUID, "originalURL", dto.OriginalURL)

	return dto.OriginalURL, nil
}

func makeSimpleUIDString(uidLength int) (string, error) {
	b := make([]byte, uidLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}
