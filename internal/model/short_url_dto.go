package model

// ShortURLDto represents a data transfer object for short URL information.
// It contains the short URL identifier, the original URL, and a deletion flag.
type ShortURLDto struct {
	ShortUID    string // The unique identifier for the short URL
	OriginalURL string // The original URL that was shortened
	IsDeleted   bool   // Indicates whether the short URL has been marked as deleted
}

// NewShortURLDto creates a new ShortURLDto instance with the provided values.
// Parameters:
//   - shortUID: the unique identifier for the short URL
//   - originalURL: the original URL that was shortened
//   - isDeleted: whether the short URL is marked as deleted
//
// Returns a new ShortURLDto instance.
func NewShortURLDto(shortUID string, originalURL string, isDeleted bool) ShortURLDto {
	return ShortURLDto{
		ShortUID:    shortUID,
		OriginalURL: originalURL,
		IsDeleted:   isDeleted,
	}
}
