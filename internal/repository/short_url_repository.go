package repository

import "github.com/Sorrowful-free/short-url-service/internal/model"

type ShortURLRepository interface {
	Save(shortURL model.ShortURLDto) error
	ContainsUID(shortUID string) bool
	GetByUID(shortUID string) (model.ShortURLDto, error)
}
