package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStatServiceImp_getStats(t *testing.T) {
	ctx := context.Background()

	t.Run("positive case - get stats successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockShortURLRepository(ctrl)
		service := NewStatService(mockRepo)

		expectedStats := model.StatDto{
			Urls:  100,
			Users: 50,
		}

		mockRepo.EXPECT().
			GetStats(ctx).
			Return(expectedStats, nil).
			Times(1)

		stats, err := service.GetStats(ctx)

		assert.NoError(t, err, "getStats should not return error")
		assert.Equal(t, expectedStats, stats, "stats should match expected values")
		assert.Equal(t, 100, stats.Urls, "urls count should be 100")
		assert.Equal(t, 50, stats.Users, "users count should be 50")
	})

	t.Run("negative case - repository returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockShortURLRepository(ctrl)
		service := NewStatService(mockRepo)

		expectedError := errors.New("database connection error")

		mockRepo.EXPECT().
			GetStats(ctx).
			Return(model.StatDto{}, expectedError).
			Times(1)

		stats, err := service.GetStats(ctx)

		assert.Error(t, err, "getStats should return error")
		assert.Equal(t, expectedError, err, "error should match expected error")
		assert.Equal(t, model.StatDto{}, stats, "stats should be empty when error occurs")
	})

	t.Run("case with zero stats", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockShortURLRepository(ctrl)
		service := NewStatService(mockRepo)

		expectedStats := model.StatDto{
			Urls:  0,
			Users: 0,
		}

		mockRepo.EXPECT().
			GetStats(ctx).
			Return(expectedStats, nil).
			Times(1)

		stats, err := service.GetStats(ctx)

		assert.NoError(t, err, "getStats should not return error")
		assert.Equal(t, expectedStats, stats, "stats should match expected values")
		assert.Equal(t, 0, stats.Urls, "urls count should be 0")
		assert.Equal(t, 0, stats.Users, "users count should be 0")
	})
}
