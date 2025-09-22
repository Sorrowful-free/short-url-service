package crypto

type UserIDEncryptor interface {
	Encrypt(userID string) (string, error)
	Decrypt(encryptedUserID string) (string, error)
}
