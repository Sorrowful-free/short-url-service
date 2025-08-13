package service

type ShortURLService interface {
	TryMakeShort(originalURL string) (string, error)
	TryMakeOriginal(shortURL string) (string, error)
}
