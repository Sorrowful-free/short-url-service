package crypto

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomSequence(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomSequenceString(size int) (string, error) {
	b, err := GenerateRandomSequence(size)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b), nil
}
