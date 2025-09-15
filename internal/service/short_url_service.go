package service

import "context"

type ShortURLService interface {
	TryMakeShort(ctx context.Context, originalURL string) (string, error)
	TryMakeOriginal(ctx context.Context, shortURL string) (string, error)
}
