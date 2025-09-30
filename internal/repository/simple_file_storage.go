package repository

import (
	"encoding/json"
	"os"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type SimpleFileStorage struct {
	fileStoragePath string
}

func NewSimpleFileStorage(fileStoragePath string) *SimpleFileStorage {
	return &SimpleFileStorage{
		fileStoragePath: fileStoragePath,
	}
}

func (sfs *SimpleFileStorage) SaveAll(userShortURLs map[string][]model.ShortURLSafeDto) error {
	jsonFile, err := os.Create(sfs.fileStoragePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	err = json.NewEncoder(jsonFile).Encode(userShortURLs)
	return err
}

func (sfs *SimpleFileStorage) LoadAll() (map[string][]model.ShortURLSafeDto, error) {
	userShortURLs := make(map[string][]model.ShortURLSafeDto)

	jsonFile, err := os.Open(sfs.fileStoragePath)

	//if file doesn't exist we must run app anyway but return empty list
	if os.IsNotExist(err) {
		return userShortURLs, nil
	} else if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&userShortURLs)
	if err != nil {
		return nil, err
	}
	return userShortURLs, nil
}
