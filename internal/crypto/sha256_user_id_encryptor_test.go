package crypto

import (
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestSha256UserIDEncryptor(t *testing.T) {
	t.Run("positive case encrypt and decrypt user ID", func(t *testing.T) {

		encryptor, err := NewSha256UserIDEncryptor(consts.TestUserIDCriptoKey)
		if err != nil {
			t.Fatalf("failed to create encryptor: %v", err)
		}

		encryptedUserID, err := encryptor.Encrypt(consts.TestUserID)
		if err != nil {
			t.Fatalf("failed to encrypt user ID: %v", err)
		}
		decryptedUserID, err := encryptor.Decrypt(encryptedUserID)
		if err != nil {
			t.Fatalf("failed to decrypt user ID: %v", err)
		}
		assert.Equal(t, consts.TestUserID, decryptedUserID, "decrypted user ID must be the same as the original user ID")
	})

	t.Run("negative case decrypt user ID", func(t *testing.T) {
		encryptor, err := NewSha256UserIDEncryptor(consts.TestUserIDCriptoKey)
		if err != nil {
			t.Fatalf("failed to create encryptor: %v", err)
		}
		_, err = encryptor.Decrypt(consts.TestInvalidUserID)
		assert.Error(t, err, "decrypt user ID must return an error")
	})
}
