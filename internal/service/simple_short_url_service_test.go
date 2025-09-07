package service

import (
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/stretchr/testify/assert"
)

func TestFService(t *testing.T) {

	l, err := logger.NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	service := NewSimpleService(consts.TestUIDLength, consts.TestFileStoragePath, l)
	originalURL := consts.TestOriginalURL
	shortUID := consts.TestShortUID

	t.Run("constructor for fake service", func(t *testing.T) {

		assert.NotNil(t, service.shortUIDs, "inner map for short urls must be not nil")
	})

	t.Run("generation of fake uid", func(t *testing.T) {

		uid, err := makeSimpleUIDString(consts.TestUIDLength)
		assert.NotEmpty(t, uid, "generation of uid must generate some string")
		assert.NoError(t, err, "generation of uid must complete without any error")
	})

	t.Run("trying to make short url", func(t *testing.T) {

		tmpShortUID, err := service.TryMakeShort(originalURL)
		assert.NotEmpty(t, tmpShortUID, "short url must be not empty")
		assert.NoError(t, err, "short url must generate without any error")
		shortUID = tmpShortUID
	})

	t.Run("trying to make original url", func(t *testing.T) {

		tmpOriginalURL, err := service.TryMakeOriginal(shortUID)
		assert.NotEmpty(t, tmpOriginalURL, "original url must be not empty")
		assert.NoError(t, err, "original url must generate without any error")
		assert.Equal(t, originalURL, tmpOriginalURL, "service must return the same url")
	})
}
