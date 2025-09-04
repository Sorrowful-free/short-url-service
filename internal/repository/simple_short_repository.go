package repository

import (
	"encoding/json"
	"os"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type SimpleShortURLRepository struct {
	shortURLs       []model.ShortURLSafeDto
	fileStoragePath string
}

func NewSimpleShortURLRepository(fileStoragePath string) *SimpleShortURLRepository {
	if fileStoragePath != "" {
		shortURLs, err := loadFromFile(fileStoragePath)
		if err != nil {
			return &SimpleShortURLRepository{
				shortURLs:       shortURLs,
				fileStoragePath: fileStoragePath,
			}
		}
	}
	return &SimpleShortURLRepository{
		shortURLs: make([]model.ShortURLSafeDto, 0),
	}
}

func (r *SimpleShortURLRepository) Save(shortURL model.ShortURLDto) error {
	r.shortURLs = append(r.shortURLs, model.NewShortURLSafeDto(shortURL))
	r.safeToFile()
	return nil
}

func (r *SimpleShortURLRepository) Load() ([]model.ShortURLSafeDto, error) {
	return r.shortURLs, nil
}

func (r *SimpleShortURLRepository) safeToFile() error {
	jsonFile, err := os.Create(r.fileStoragePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	json.NewEncoder(jsonFile).Encode(r.shortURLs)
	return nil
}

func loadFromFile(fileStoragePath string) ([]model.ShortURLSafeDto, error) {
	shortURLs := make([]model.ShortURLSafeDto, 0)
	jsonFile, err := os.Open(fileStoragePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	json.NewDecoder(jsonFile).Decode(&shortURLs)
	return shortURLs, nil
}
