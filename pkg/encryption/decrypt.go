package encryption

import (
	"crypto/cipher"
)

func (cipherBlock CipherBlock) DecryptData(encryptedData []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(cipherBlock.Block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
