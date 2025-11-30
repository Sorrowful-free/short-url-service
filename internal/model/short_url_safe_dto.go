package model

import "github.com/google/uuid"

// ShortURLSafeDto represents a short URL data transfer object with a unique UUID.
// It is used for safe storage and serialization of short URL data.
// The UUID field provides an additional unique identifier for the record.
type ShortURLSafeDto struct {
	UUID        string `json:"uuid"`         // Unique identifier for the record
	ShortUID    string `json:"short_url"`    // The short URL identifier (note: naming kept for autotest compatibility)
	OriginalURL string `json:"original_url"` // The original URL that was shortened
	IsDeleted   bool   `json:"is_deleted"`   // Indicates whether the short URL has been marked as deleted
}

// NewShortURLSafeDto creates a new ShortURLSafeDto from a ShortURLDto.
// It generates a new UUID for the safe DTO.
// Parameters:
//   - ShortURLDto: the source ShortURLDto to convert
//
// Returns a new ShortURLSafeDto instance with a generated UUID.
func NewShortURLSafeDto(ShortURLDto ShortURLDto) ShortURLSafeDto {
	return ShortURLSafeDto{
		UUID:        uuid.New().String(),
		ShortUID:    ShortURLDto.ShortUID,
		OriginalURL: ShortURLDto.OriginalURL,
		IsDeleted:   ShortURLDto.IsDeleted,
	}
}
