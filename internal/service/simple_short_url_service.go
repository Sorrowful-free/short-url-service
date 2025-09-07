package service

import (
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/repository"
)

type SimpleShortURLService struct {
	shortUIDs          map[string]model.ShortURLDto
	uidLength          int
	logger             *logger.Logger
	ShortURLRepository repository.ShortURLRepository
}

func NewSimpleService(uidLength int, fileStoragePath string, logger *logger.Logger) *SimpleShortURLService {
	shortURLRepository := repository.NewSimpleShortURLRepository(fileStoragePath)
	service := SimpleShortURLService{
		shortUIDs:          make(map[string]model.ShortURLDto),
		uidLength:          uidLength,
		logger:             logger,
		ShortURLRepository: shortURLRepository,
	}
	service.restoreShortURLsFromRepository()
	return &service
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

	service.ShortURLRepository.Save(dto)

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

func (service SimpleShortURLService) restoreShortURLsFromRepository() error {
	shortURLs, err := service.ShortURLRepository.Load()
	if err != nil {
		return err
	}
	service.shortUIDs = make(map[string]model.ShortURLDto)
	for _, shortSafeURL := range shortURLs {
		service.shortUIDs[shortSafeURL.ShortUID] = model.New(shortSafeURL.ShortUID, shortSafeURL.OriginalURL)
	}
	return nil
}
