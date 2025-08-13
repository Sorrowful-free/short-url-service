package service

import (
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type FakeShortUrlService struct {
	shortUrls    map[string]model.ShortUrlDto
	originalUrls map[string]model.ShortUrlDto
}

func NewFakeService() *FakeShortUrlService {
	return &FakeShortUrlService{
		shortUrls:    make(map[string]model.ShortUrlDto),
		originalUrls: make(map[string]model.ShortUrlDto),
	}
}

func (service FakeShortUrlService) TryMakeShort(originalUrl string) (string, error) {

	_, exist := service.originalUrls[originalUrl]
	if exist {
		return "", fmt.Errorf("url %s already exist ", originalUrl)
	}

	shortUrl, err := makeFakeUIDString()
	if err != nil {
		return shortUrl, err
	}
	dto := model.New(shortUrl, originalUrl)

	service.shortUrls[shortUrl] = dto
	service.originalUrls[originalUrl] = dto

	return shortUrl, nil
}

func (service FakeShortUrlService) TryMakeOriginal(shortUrl string) (string, error) {
	dto, exist := service.shortUrls[shortUrl]

	if !exist {
		return "", fmt.Errorf("short url %s doesnot exist ", shortUrl)
	}

	return dto.OriginalUrl, nil
}

func makeFakeUIDString() (string, error) {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}
