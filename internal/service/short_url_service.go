package service

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type ShortURLService interface {
	TryMakeShort(ctx context.Context, userID string, originalURL string) (string, error)
	TryMakeOriginal(ctx context.Context, shortURL string) (string, error)
	TryMakeShortBatch(ctx context.Context, userID string, originalURLs []string) ([]string, error)
	GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error)
	Ping(ctx context.Context) error
}
