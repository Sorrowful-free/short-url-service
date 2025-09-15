package repository

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type ShortURLRepository interface {
	Save(ctx context.Context, shortURL model.ShortURLDto) error
	ContainsUID(ctx context.Context, shortUID string) bool
	GetByUID(ctx context.Context, shortUID string) (model.ShortURLDto, error)
	Ping(ctx context.Context) error
}
