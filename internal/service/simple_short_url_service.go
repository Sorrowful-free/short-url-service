package service

import (
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type SimpleShortURLService struct {
	shortUIDs map[string]model.ShortURLDto
	uidLength int
}

func NewSimpleService(uidLength int) *SimpleShortURLService {
	return &SimpleShortURLService{
		shortUIDs: make(map[string]model.ShortURLDto),
		uidLength: uidLength,
	}
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

	return shortUID, nil
}

func (service SimpleShortURLService) TryMakeOriginal(shortUID string) (string, error) {
	dto, exist := service.shortUIDs[shortUID]

	if !exist {
		return "", fmt.Errorf("short url %s doesnot exist ", shortUID)
	}

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
