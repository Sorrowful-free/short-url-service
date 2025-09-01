package service

import (
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type FakeShortURLService struct {
	baseURL      string
	shortUIDs    map[string]model.ShortURLDto
	originalURLs map[string]model.ShortURLDto
}

func NewFakeService() *FakeShortURLService {
	return &FakeShortURLService{
		shortUIDs:    make(map[string]model.ShortURLDto),
		originalURLs: make(map[string]model.ShortURLDto),
	}
}

func (service FakeShortURLService) TryMakeShort(originalURL string) (string, error) {

	_, exist := service.originalURLs[originalURL]
	if exist {
		return "", fmt.Errorf("url %s already exist ", originalURL)
	}

	shortUID, err := makeFakeUIDString()

	if err != nil {
		return shortUID, err
	}

	dto := model.New(shortUID, originalURL)

	service.shortUIDs[shortUID] = dto
	service.originalURLs[originalURL] = dto

	return shortUID, nil
}

func (service FakeShortURLService) TryMakeOriginal(shortUID string) (string, error) {
	dto, exist := service.shortUIDs[shortUID]

	if !exist {
		return "", fmt.Errorf("short url %s doesnot exist ", shortUID)
	}

	return dto.OriginalURL, nil
}

func makeFakeUIDString() (string, error) {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}
