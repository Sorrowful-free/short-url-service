package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeService(t *testing.T) {

	address := "localhost:8080"
	t.Run("constructor for fake service", func(t *testing.T) {
		service := NewFakeService(address)

		assert.NotNil(t, service.originalURLs, "inner map for original urls must be not nil")
		assert.NotNil(t, service.shortURLs, "inner map for short urls must be not nil")
	})

	t.Run("generation of fake uid", func(t *testing.T) {
		uid, err := makeFakeUIDString()
		assert.NotEmpty(t, uid, "generation of uid must generate some string")
		assert.Empty(t, err, "generation of uid must complete without any error")
	})

	t.Run("trying to make short url", func(t *testing.T) {
		service := NewFakeService(address)
		originalURL := "http://google.com"
		short, err := service.TryMakeShort(originalURL)
		assert.NotEmpty(t, short, "short url must be not empty")
		assert.Empty(t, err, "short url must generate without any error")
		_, err = service.TryMakeShort(originalURL)
		assert.NotEmpty(t, err, "short url must generate error for dublicates")
	})

	t.Run("trying to make original url", func(t *testing.T) {
		service := NewFakeService(address)
		originalURL := "http://google.com"
		short, err := service.TryMakeShort(originalURL)
		assert.NotEmpty(t, short, "short url must be not empty")
		assert.Empty(t, err, "short url must generate without any error")
		otherOriginalURL, err := service.TryMakeOriginal(short)
		assert.Empty(t, err, "original url must returned without any error")
		assert.Equal(t, originalURL, otherOriginalURL, "service must return the same url")
	})
}
