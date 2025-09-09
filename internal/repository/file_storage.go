package repository

import "github.com/Sorrowful-free/short-url-service/internal/model"

type FileStorage interface {
	SafeAll(shortURLs []model.ShortURLSafeDto) error
	LoadAll() ([]model.ShortURLSafeDto, error)
}
