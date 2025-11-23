package service

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

//go:generate mockgen -source=short_url_service.go -destination=./../../mocks/mock_short_url_service.go --package=mocks
type ShortURLService interface {
	TryMakeShort(ctx context.Context, userID string, originalURL string) (model.ShortURLDto, error)
	TryMakeOriginal(ctx context.Context, shortURL string) (model.ShortURLDto, error)
	TryMakeShortBatch(ctx context.Context, userID string, originalURLs []string) ([]model.ShortURLDto, error)
	GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error)
	DeleteShortURLs(ctx context.Context, userID string, shortURLs []string) error
	Ping(ctx context.Context) error
}
