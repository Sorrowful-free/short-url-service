package service

import "github.com/Sorrowful-free/short-url-service/internal/model"

type FakeShortUrlService struct {
	urls map[string]model.ShortUrlDto
}

func NewFakeService() *FakeShortUrlService {
	return &FakeShortUrlService{
		urls: make(map[string]model.ShortUrlDto),
	}
}

func (service FakeShortUrlService) TryMakeShort(originalUrl string) (string, *error) {

	return "short_url", nil
}

func (service FakeShortUrlService) TryMakeOriginal(shortUrl string) (string, *error) {
	return "long_url", nil
}
