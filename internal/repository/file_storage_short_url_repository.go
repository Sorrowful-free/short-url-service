package repository

import (
	"context"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type FileStorageShortURLRepository struct {
	shortURLs   []model.ShortURLSafeDto
	fileStorage FileStorage
}

func NewFileStorageShortURLRepository(fileStoragePath string) (ShortURLRepository, error) {
	fileStorage := NewSimpleFileStorage(fileStoragePath)
	shortURLs, err := fileStorage.LoadAll()
	if err != nil {
		return nil, err
	}
	return &FileStorageShortURLRepository{
		shortURLs:   shortURLs,
		fileStorage: fileStorage,
	}, nil
}

func (r *FileStorageShortURLRepository) Save(ctx context.Context, shortURL model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	r.shortURLs = append(r.shortURLs, model.NewShortURLSafeDto(shortURL))
	r.fileStorage.SafeAll(r.shortURLs)
	return nil
}

func (r *FileStorageShortURLRepository) SaveBatch(ctx context.Context, shortURLs []model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	for _, shortURL := range shortURLs {
		r.shortURLs = append(r.shortURLs, model.NewShortURLSafeDto(shortURL))
	}
	r.fileStorage.SafeAll(r.shortURLs)
	return nil
}

func (r *FileStorageShortURLRepository) ContainsUID(ctx context.Context, shortUID string) bool {
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

func (r *FileStorageShortURLRepository) GetByUID(ctx context.Context, shortUID string) (model.ShortURLDto, error) {
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

func (r *FileStorageShortURLRepository) GetByOriginalURL(ctx context.Context, originalURL string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}

	for _, shortURL := range r.shortURLs {
		if shortURL.OriginalURL == originalURL {
			return model.New(shortURL.ShortUID, shortURL.OriginalURL), nil
		}
	}
	return model.ShortURLDto{}, fmt.Errorf("original url %s not found", originalURL)
}

func (r *FileStorageShortURLRepository) Ping(ctx context.Context) error {
	return nil
}
