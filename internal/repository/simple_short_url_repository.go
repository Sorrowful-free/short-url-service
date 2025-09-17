package repository

import (
	"context"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type SimpleShortURLRepository struct {
	shortURLs []model.ShortURLSafeDto
}

func NewSimpleShortURLRepository(fileStoragePath string) (ShortURLRepository, error) {
	return &SimpleShortURLRepository{
		shortURLs: make([]model.ShortURLSafeDto, 0),
	}, nil
}

func (r *SimpleShortURLRepository) Save(ctx context.Context, shortURL model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if r.containsOriginalURL(ctx, shortURL.OriginalURL) {
		return NewOriginalURLConflictRepositoryError(shortURL.OriginalURL)
	}
	r.shortURLs = append(r.shortURLs, model.NewShortURLSafeDto(shortURL))

	return nil
}

func (r *SimpleShortURLRepository) SaveBatch(ctx context.Context, shortURLs []model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	for _, shortURL := range shortURLs {
		r.shortURLs = append(r.shortURLs, model.NewShortURLSafeDto(shortURL))
	}
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

func (r *SimpleShortURLRepository) GetByOriginalURL(ctx context.Context, originalURL string) (model.ShortURLDto, error) {
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

func (r *SimpleShortURLRepository) Ping(ctx context.Context) error {
	return nil
}

func (r *SimpleShortURLRepository) containsOriginalURL(ctx context.Context, originalURL string) bool {
	if ctx.Err() != nil {
		return false
	}
	for _, shortURL := range r.shortURLs {
		if shortURL.OriginalURL == originalURL {
			return true
		}
	}
	return false
}
