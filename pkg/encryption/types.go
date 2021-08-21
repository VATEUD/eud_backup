package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

type CipherBlock struct {
	cipher.Block
}

func New() (CipherBlock, error) {
	c, err := aes.NewCipher([]byte(os.Getenv("CIPHER_KEY")))

	if err != nil {
		return CipherBlock{}, err
	}

	return CipherBlock{c}, nil
}