package service

type ShortUrlService interface {
	TryMakeShort(originalUrl string) (string, error)
	TryMakeOriginal(shortUrl string) (string, error)
}
