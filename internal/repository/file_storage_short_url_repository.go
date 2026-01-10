package repository

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type FileStorageShortURLRepository struct {
	SimpleShortURLRepository
	fileStorage FileStorage
}

func NewFileStorageShortURLRepository(fileStoragePath string) (ShortURLRepository, error) {
	fileStorage := NewSimpleFileStorage(fileStoragePath)
	userShortURLs, err := fileStorage.LoadAll()
	if err != nil {
		return nil, err
	}
	return &FileStorageShortURLRepository{
		SimpleShortURLRepository: SimpleShortURLRepository{
			userShortURLs: userShortURLs,
		},
		fileStorage: fileStorage,
	}, nil
}

func (r *FileStorageShortURLRepository) Save(ctx context.Context, userID string, shortURL model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	r.SimpleShortURLRepository.Save(ctx, userID, shortURL)
	r.fileStorage.SaveAll(r.userShortURLs)
	return nil
}

func (r *FileStorageShortURLRepository) SaveBatch(ctx context.Context, userID string, shortURLs []model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	r.SimpleShortURLRepository.SaveBatch(ctx, userID, shortURLs)
	r.fileStorage.SaveAll(r.userShortURLs)
	return nil
}

func (r *FileStorageShortURLRepository) ContainsUID(ctx context.Context, shortUID string) bool {
	if ctx.Err() != nil {
		return false
	}
	return r.SimpleShortURLRepository.ContainsUID(ctx, shortUID)
}

func (r *FileStorageShortURLRepository) GetByUID(ctx context.Context, shortUID string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}
	return r.SimpleShortURLRepository.GetByUID(ctx, shortUID)
}

func (r *FileStorageShortURLRepository) GetByOriginalURL(ctx context.Context, originalURL string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}

	return r.SimpleShortURLRepository.GetByOriginalURL(ctx, originalURL)
}

func (r *FileStorageShortURLRepository) Ping(ctx context.Context) error {
	return r.SimpleShortURLRepository.Ping(ctx)
}

func (r *FileStorageShortURLRepository) GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return r.SimpleShortURLRepository.GetUserUrls(ctx, userID)
}

func (r *FileStorageShortURLRepository) GetStats(ctx context.Context) (model.StatDto, error) {
	return r.SimpleShortURLRepository.GetStats(ctx)
}
