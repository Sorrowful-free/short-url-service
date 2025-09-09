package repository

import (
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type SimpleShortURLRepository struct {
	shortURLs   []model.ShortURLSafeDto
	fileStorage FileStorage
}

func NewSimpleShortURLRepository(fileStoragePath string) (ShortURLRepository, error) {
	fileStorage := NewSimpleFileStorage(fileStoragePath)
	shortURLs, err := fileStorage.LoadAll()
	if err != nil {
		return nil, err
	}
	return &SimpleShortURLRepository{
		shortURLs:   shortURLs,
		fileStorage: fileStorage,
	}, nil
}

func (r *SimpleShortURLRepository) Save(shortURL model.ShortURLDto) error {
	r.shortURLs = append(r.shortURLs, model.NewShortURLSafeDto(shortURL))
	r.fileStorage.SafeAll(r.shortURLs)
	return nil
}

func (r *SimpleShortURLRepository) ContainsUID(shortUID string) bool {
	for _, shortURL := range r.shortURLs {
		if shortURL.ShortUID == shortUID {
			return true
		}
	}
	return false
}

func (r *SimpleShortURLRepository) GetByUID(shortUID string) (model.ShortURLDto, error) {
	for _, shortURL := range r.shortURLs {
		if shortURL.ShortUID == shortUID {
			return model.New(shortURL.ShortUID, shortURL.OriginalURL), nil
		}
	}
	return model.ShortURLDto{}, fmt.Errorf("short url %s not found", shortUID)
}
