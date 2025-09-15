package repository

import (
	"context"
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

func (r *SimpleShortURLRepository) Save(ctx context.Context, shortURL model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	r.shortURLs = append(r.shortURLs, model.NewShortURLSafeDto(shortURL))
	r.fileStorage.SafeAll(r.shortURLs)
	return nil
}

func (r *SimpleShortURLRepository) ContainsUID(ctx context.Context, shortUID string) bool {
	if ctx.Err() != nil {
		return false
	}
	for _, shortURL := range r.shortURLs {
		if shortURL.ShortUID == shortUID {
			return true
		}
	}
	return false
}

func (r *SimpleShortURLRepository) GetByUID(ctx context.Context, shortUID string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}
	for _, shortURL := range r.shortURLs {
		if shortURL.ShortUID == shortUID {
			return model.New(shortURL.ShortUID, shortURL.OriginalURL), nil
		}
	}
	return model.ShortURLDto{}, fmt.Errorf("short url %s not found", shortUID)
}

func (r *SimpleShortURLRepository) Ping(ctx context.Context) error {
	return nil
}
