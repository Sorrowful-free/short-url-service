package crypto

// UserIDEncryptor defines the interface for encrypting and decrypting user IDs.
// It is used to securely store user IDs in cookies.
type UserIDEncryptor interface {
	// Encrypt encrypts a user ID string.
	// Returns the encrypted string and an error if encryption fails.
	Encrypt(userID string) (string, error)

	// Decrypt decrypts an encrypted user ID string.
	// Returns the decrypted user ID and an error if decryption fails.
	Decrypt(encryptedUserID string) (string, error)
}
