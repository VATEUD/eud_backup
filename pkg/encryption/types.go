package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

// CipherBlock represents the cipher that'll be used to encrypt bytes
type CipherBlock struct {
	cipher.Block
}

// New constructs a new cipher that'll be used for data encryption
func New() (CipherBlock, error) {
	c, err := aes.NewCipher([]byte(os.Getenv("CIPHER_KEY")))

	if err != nil {
		return CipherBlock{}, err
	}

	return CipherBlock{c}, nil
}
