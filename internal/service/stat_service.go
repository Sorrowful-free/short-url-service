package service

import (
	"context"

	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/repository"
)

//go:generate mockgen -source=stat_service.go -destination=./../../mocks/mock_stat_service.go --package=mocks
type StatService interface {
	GetStats(ctx context.Context) (model.StatDto, error)
}

type StatServiceImp struct {
	shortURLRepository repository.ShortURLRepository
}

func NewStatService(shortURLRepository repository.ShortURLRepository) StatService {
	s := &StatServiceImp{
		shortURLRepository: shortURLRepository,
	}

	return s
}

func (s *StatServiceImp) GetStats(ctx context.Context) (model.StatDto, error) {
	return s.shortURLRepository.GetStats(ctx)
}
