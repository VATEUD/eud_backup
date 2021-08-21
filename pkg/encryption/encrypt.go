package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
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

func (cipherBlock CipherBlock) DecryptData(encryptedData []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(cipherBlock.Block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
