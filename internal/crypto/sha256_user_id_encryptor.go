package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
)

type Sha256UserIDEncryptor struct {
	userIDCriptoKeyHash [32]byte
	aesblock            cipher.Block
	aesgcm              cipher.AEAD
	nonce               []byte
}

func NewSha256UserIDEncryptor(userIDCriptoKey string) (UserIDEncryptor, error) {

	key := getUserIDCriptoKeyHash(userIDCriptoKey)

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	nonce, err := GenerateRandomSequence(aesgcm.NonceSize())
	if err != nil {
		return nil, err
	}

	return &Sha256UserIDEncryptor{
		userIDCriptoKeyHash: getUserIDCriptoKeyHash(userIDCriptoKey),
		aesblock:            aesblock,
		aesgcm:              aesgcm,
		nonce:               nonce,
	}, nil
}

func (e *Sha256UserIDEncryptor) Encrypt(userID string) (string, error) {
	userIDBytes := e.aesgcm.Seal(nil, e.nonce, []byte(userID), nil)
	userIDString := base64.StdEncoding.EncodeToString(userIDBytes)
	return userIDString, nil
}

func (e *Sha256UserIDEncryptor) Decrypt(encryptedUserID string) (string, error) {
	userIDBytes, err := base64.StdEncoding.DecodeString(encryptedUserID)
	if err != nil {
		return "", err
	}
	decryptedUserID, err := e.aesgcm.Open(nil, e.nonce, userIDBytes, nil)
	if err != nil {
		return "", err
	}
	return string(decryptedUserID), nil
}

func getUserIDCriptoKeyHash(userIDCriptoKey string) [32]byte {

	return sha256.Sum256([]byte(userIDCriptoKey))
}
