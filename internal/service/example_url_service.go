package service

import (
	"context"
	"errors"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type ExampleURLService struct {
	ConflictURL string
	HasURLs     bool
	PingError   bool
}

func (s *ExampleURLService) SetConflictURL(conflictURL string) *ExampleURLService {
	s.ConflictURL = conflictURL
	return s
}

func (s *ExampleURLService) SetHasURLs(hasURLs bool) *ExampleURLService {
	s.HasURLs = hasURLs
	return s
}

func (s *ExampleURLService) SetPingError(pingError bool) *ExampleURLService {
	s.PingError = pingError
	return s
}

func (s *ExampleURLService) TryMakeShort(ctx context.Context, userID string, originalURL string) (model.ShortURLDto, error) {
	if originalURL == s.ConflictURL {
		return model.NewShortURLDto("abc123", originalURL, false), NewOriginalURLConflictServiceError(originalURL)
	}
	return model.NewShortURLDto("abc123", originalURL, false), nil
}

func (s *ExampleURLService) TryMakeOriginal(ctx context.Context, shortURL string) (model.ShortURLDto, error) {
	if shortURL == "deleted123" {
		return model.ShortURLDto{ShortUID: shortURL, OriginalURL: "https://example.com/original-url", IsDeleted: true}, nil
	}
	return model.ShortURLDto{ShortUID: shortURL, OriginalURL: "https://example.com/original-url", IsDeleted: false}, nil
}

func (s *ExampleURLService) TryMakeShortBatch(ctx context.Context, userID string, originalURLs []string) ([]model.ShortURLDto, error) {
	result := make([]model.ShortURLDto, len(originalURLs))
	shortUIDs := []string{"abc123", "def456", "ghi789"}
	for i, url := range originalURLs {
		if i < len(shortUIDs) {
			result[i] = model.NewShortURLDto(shortUIDs[i], url, false)
		}
	}
	return result, nil
}

func (s *ExampleURLService) GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error) {
	if !s.HasURLs {
		return []model.ShortURLDto{}, nil
	}
	return []model.ShortURLDto{
		{ShortUID: "abc123", OriginalURL: "https://example.com/url1", IsDeleted: false},
		{ShortUID: "def456", OriginalURL: "https://example.com/url2", IsDeleted: false},
	}, nil
}

func (s *ExampleURLService) DeleteShortURLs(ctx context.Context, userID string, shortURLs []string) error {
	return nil
}

func (s *ExampleURLService) Ping(ctx context.Context) error {
	if s.PingError {
		return errors.New("database connection error")
	}
	return nil
}
