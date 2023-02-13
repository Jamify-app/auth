package encryption

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/scrypt"
)

func EncryptPassword(password string) ([]byte, []byte, error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, nil, err
	}

	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, 64)
	if err != nil {
		return nil, nil, err
	}

	return salt, hash, nil
}
