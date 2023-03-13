package encryption_tests

import (
	"testing"

	"github.com/Jamify-app/auth/common/encryption"
	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	t.Run("Should encrypt password and return salt, hash, and no error", func(t *testing.T) {
		salt, hash, err := encryption.EncryptPassword("password1234!")
		assert.NotEmpty(t, salt)
		assert.NotEmpty(t, hash)
		assert.Empty(t, err)
		assert.Equal(t, 64, len(hash))
		assert.Equal(t, 32, len(salt))
	})

	t.Run("Should return no salt, no hash, and an error for no password", func(t *testing.T) {
		salt, hash, err := encryption.EncryptPassword("")
		assert.Empty(t, salt)
		assert.Empty(t, hash)
		assert.Error(t, err)
	})
}
