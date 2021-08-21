package encryption

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encrypt encrypts the given bytes
func (cipherBlock CipherBlock) Encrypt(file []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(cipherBlock.Block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, file, nil), nil
}
