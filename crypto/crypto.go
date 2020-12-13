package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encrypt takes in original data and key and
// returns encrypted data
func Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encdata := gcm.Seal(nonce, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	return encdata, nil
}

// Decrypt Takes in encrypted data and key and
// returns decrypted data
func Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := data[:gcm.NonceSize()]
	encdata := data[gcm.NonceSize():]
	decdata, err := gcm.Open(nil, nonce, encdata, nil)
	if err != nil {
		return nil, err
	}
	return decdata, nil
}
