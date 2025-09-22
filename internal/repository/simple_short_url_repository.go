package repository

import (
	"context"
	"fmt"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type SimpleShortURLRepository struct {
	userShortURLs map[string][]model.ShortURLSafeDto
}

func NewSimpleShortURLRepository(fileStoragePath string) (ShortURLRepository, error) {
	return &SimpleShortURLRepository{
		userShortURLs: make(map[string][]model.ShortURLSafeDto),
	}, nil
}

func (r *SimpleShortURLRepository) Save(ctx context.Context, userID string, shortURL model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if _, ok := r.userShortURLs[userID]; !ok {
		r.userShortURLs[userID] = make([]model.ShortURLSafeDto, 0)
	}
	r.userShortURLs[userID] = append(r.userShortURLs[userID], model.NewShortURLSafeDto(shortURL))

	return nil
}

func (r *SimpleShortURLRepository) SaveBatch(ctx context.Context, userID string, shortURLs []model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	for _, shortURL := range shortURLs {
		err := r.Save(ctx, userID, shortURL)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *SimpleShortURLRepository) ContainsUID(ctx context.Context, shortUID string) bool {
	if ctx.Err() != nil {
		return false
	}

	for _, shortURLs := range r.userShortURLs {
		for _, shortURL := range shortURLs {
			if shortURL.ShortUID == shortUID {
				return true
			}
		}
	}

	return false
}

func (r *SimpleShortURLRepository) GetByUID(ctx context.Context, shortUID string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}

	for _, shortURLs := range r.userShortURLs {
		for _, shortURL := range shortURLs {
			if shortURL.ShortUID == shortUID {
				return model.NewShortURLDto(shortURL.ShortUID, shortURL.OriginalURL), nil
			}
		}
	}
	return model.ShortURLDto{}, fmt.Errorf("short url %s not found", shortUID)
}

func (r *SimpleShortURLRepository) GetByOriginalURL(ctx context.Context, originalURL string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}

	for _, shortURLs := range r.userShortURLs {
		for _, shortURL := range shortURLs {
			if shortURL.OriginalURL == originalURL {
				return model.NewShortURLDto(shortURL.ShortUID, shortURL.OriginalURL), nil
			}
		}
	}
	return model.ShortURLDto{}, fmt.Errorf("original url %s not found", originalURL)
}

func (r *SimpleShortURLRepository) Ping(ctx context.Context) error {
	return nil
}

func (r *SimpleShortURLRepository) GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	shortURLs, ok := r.userShortURLs[userID]

	if !ok {
		return nil, fmt.Errorf("user %s not found", userID)
	}

	shortURLDtos := make([]model.ShortURLDto, len(shortURLs))
	for i, shortURL := range shortURLs {
		shortURLDtos[i] = model.NewShortURLDto(shortURL.ShortUID, shortURL.OriginalURL)
	}

	return shortURLDtos, nil
}
