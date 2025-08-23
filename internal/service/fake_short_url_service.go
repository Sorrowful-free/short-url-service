package service

import (
	"crypto/rand"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type FakeShortURLService struct {
	baseURL   string
	shortUIDs map[string]model.ShortURLDto
	uidLength int
}

func NewFakeService(uidLength int) *FakeShortURLService {
	return &FakeShortURLService{
		shortUIDs: make(map[string]model.ShortURLDto),
		uidLength: uidLength,
	}
}

func (service FakeShortURLService) TryMakeShort(originalURL string) (string, error) {

	shortUID, err := makeFakeUIDString(service.uidLength)

	if err != nil {
		return shortUID, err
	}

	dto := model.New(shortUID, originalURL)

	service.shortUIDs[shortUID] = dto

	return shortUID, nil
}

func (service FakeShortURLService) TryMakeOriginal(shortUID string) (string, error) {
	dto, exist := service.shortUIDs[shortUID]

	if !exist {
		return "", fmt.Errorf("short url %s doesnot exist ", shortUID)
	}

	return dto.OriginalURL, nil
}

func makeFakeUIDString(uidLength int) (string, error) {
	b := make([]byte, uidLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}
