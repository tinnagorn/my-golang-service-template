package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/pkg/errors"
)

func DecryptAESGCM(src string, key []byte, iv []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", errors.Wrap(err, "decode base64")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.Wrap(err, "new cipher")
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "new GCM")
	}

	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", errors.Wrap(err, "open")
	}

	return string(plaintext), nil
}
