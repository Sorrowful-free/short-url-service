package service

import (
	"context"
	"errors"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type ExampleStatService struct {
	GetStatsError bool
}

func (s *ExampleStatService) SetGetStatsError(getStatsError bool) *ExampleStatService {
	s.GetStatsError = getStatsError
	return s
}

func (s *ExampleStatService) GetStats(ctx context.Context) (model.StatDto, error) {
	if s.GetStatsError {
		return model.StatDto{}, errors.New("database connection error")
	}
	return model.StatDto{Urls: 1, Users: 1}, nil
}
