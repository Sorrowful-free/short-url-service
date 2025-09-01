package service

import (
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type FakeShortURLService struct {
	baseURL      string
	shortURLs    map[string]model.ShortURLDto
	originalURLs map[string]model.ShortURLDto
}

func NewFakeService(baseURL string) *FakeShortURLService {
	return &FakeShortURLService{
		baseURL:      baseURL,
		shortURLs:    make(map[string]model.ShortURLDto),
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

	shortURL := fmt.Sprintf("%s://%s/%s", "http", service.baseURL, shortUID)
	dto := model.New(shortURL, originalURL)

	service.shortURLs[shortUID] = dto
	service.originalURLs[originalURL] = dto

	return shortURL, nil
}

func (service FakeShortURLService) TryMakeOriginal(shortUID string) (string, error) {
	dto, exist := service.shortURLs[shortUID]

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
