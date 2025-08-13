package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeService(t *testing.T) {
	t.Run("constructor for fake service", func(t *testing.T) {
		service := NewFakeService()

		assert.NotNil(t, service.originalUrls, "inner map for original urls must be not nil")
		assert.NotNil(t, service.shortUrls, "inner map for short urls must be not nil")
	})

	t.Run("generation of fake uid", func(t *testing.T) {
		uid, err := makeFakeUIDString()
		assert.NotEmpty(t, uid, "generation of uid must generate some string")
		assert.Empty(t, err, "generation of uid must complete without any error")
	})

	t.Run("trying to make short url", func(t *testing.T) {
		service := NewFakeService()
		originalUrl := "http://google.com"
		short, err := service.TryMakeShort(originalUrl)
		assert.NotEmpty(t, short, "short url must be not empty")
		assert.Empty(t, err, "short url must generate without any error")
		_, err = service.TryMakeShort(originalUrl)
		assert.NotEmpty(t, err, "short url must generate error for dublicates")
	})

	t.Run("trying to make original url", func(t *testing.T) {
		service := NewFakeService()
		originalUrl := "http://google.com"
		short, err := service.TryMakeShort(originalUrl)
		assert.NotEmpty(t, short, "short url must be not empty")
		assert.Empty(t, err, "short url must generate without any error")
		otherOriginalUrl, err := service.TryMakeOriginal(short)
		assert.Empty(t, err, "original url must returned without any error")
		assert.Equal(t, originalUrl, otherOriginalUrl, "service must return the same url")
	})
}
