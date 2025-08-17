package service

import (
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeService(t *testing.T) {

	address := "localhost:8080"
	service := NewFakeService(address)
	originalURL := "http://google.com"
	shortURL := "http://localhost:8080/1234567890"

	t.Run("constructor for fake service", func(t *testing.T) {

		assert.NotNil(t, service.originalURLs, "inner map for original urls must be not nil")
		assert.NotNil(t, service.shortURLs, "inner map for short urls must be not nil")
	})

	t.Run("generation of fake uid", func(t *testing.T) {

		uid, err := makeFakeUIDString()
		assert.NotEmpty(t, uid, "generation of uid must generate some string")
		assert.NoError(t, err, "generation of uid must complete without any error")
	})

	t.Run("trying to make short url", func(t *testing.T) {

		tmpShortURL, err := service.TryMakeShort(originalURL)
		assert.NotEmpty(t, tmpShortURL, "short url must be not empty")
		assert.NoError(t, err, "short url must generate without any error")
		_, err = service.TryMakeShort(originalURL)
		assert.NotEmpty(t, err, "short url must generate error for dublicates")
		shortURL = tmpShortURL
	})

	t.Run("trying to make original url", func(t *testing.T) {

		u, _ := url.Parse(shortURL)
		shortUID := path.Base(u.Path)

		tmpOriginalURL, err := service.TryMakeOriginal(shortUID)
		assert.NotEmpty(t, tmpOriginalURL, "original url must be not empty")
		assert.NoError(t, err, "original url must generate without any error")
		assert.Equal(t, originalURL, tmpOriginalURL, "service must return the same url")
	})
}
