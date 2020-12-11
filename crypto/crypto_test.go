package crypto

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
)

func TestCrypto(t *testing.T) {
	datalen := 50
	keylen := 32
	data := make([]byte, datalen)
	key := make([]byte, keylen)

	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		t.Errorf("Random data generation failed")
	}
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		t.Errorf("Random key generation failed")
	}

	encdata, err := encrypt(data, key)
	if err != nil {
		t.Errorf("Encryption failed")
	}

	decdata, err := decrypt(encdata, key)
	if err != nil {
		t.Errorf("Decryption failed")
	}

	if res := bytes.Compare(data, decdata); res != 0 {
		t.Errorf("Decrypted data does not match original data")
	}
}
