package service

import (
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type FakeShortURLService struct {
	shortURLs    map[string]model.ShortURLDto
	originalURLs map[string]model.ShortURLDto
}

func NewFakeService() *FakeShortURLService {
	return &FakeShortURLService{
		shortURLs:    make(map[string]model.ShortURLDto),
		originalURLs: make(map[string]model.ShortURLDto),
	}
}

func (service FakeShortURLService) TryMakeShort(originalURL string) (string, error) {

	_, exist := service.originalURLs[originalURL]
	if exist {
		return "", fmt.Errorf("url %s already exist ", originalURL)
	}

	shortURL, err := makeFakeUIDString()
	if err != nil {
		return shortURL, err
	}
	dto := model.New(shortURL, originalURL)

	service.shortURLs[shortURL] = dto
	service.originalURLs[originalURL] = dto

	return shortURL, nil
}

func (service FakeShortURLService) TryMakeOriginal(shortURL string) (string, error) {
	dto, exist := service.shortURLs[shortURL]

	if !exist {
		return "", fmt.Errorf("short url %s doesnot exist ", shortURL)
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
