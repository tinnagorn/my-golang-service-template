package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/pkg/errors"
)

func EncryptAESGCM(src string, key []byte, iv []byte) (string, error) {
	plaintext := []byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.WithStack(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.WithStack(err)
	}

	ciphertext := aesgcm.Seal(nil, iv, plaintext, nil)
	result64 := base64.StdEncoding.EncodeToString(ciphertext)

	return result64, nil
}
