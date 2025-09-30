package repository

import "github.com/Sorrowful-free/short-url-service/internal/model"

type FileStorage interface {
	SaveAll(userShortURLs map[string][]model.ShortURLSafeDto) error
	LoadAll() (map[string][]model.ShortURLSafeDto, error)
}
