package model

type ShortUrlDto struct {
	ShortUrl    string
	OriginalUrl string
}

func New(shortUrl string, originalUrl string) ShortUrlDto {
	return ShortUrlDto{
		ShortUrl:    shortUrl,
		OriginalUrl: originalUrl,
	}
}
