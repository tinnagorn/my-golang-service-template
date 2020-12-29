package utility

import "testing"

func TestEncryptThenDecryptShouldGetTheOldValue(t *testing.T) {
	testMsg := "This value should be the same values + ภาษาไทย และตัวเลข 0123456789 ด้วย"
	key := []byte("59azblwa8d64f6a718p061156a21oi68")
	iv := make([]byte, 12)
	encryptedVal, err := EncryptAESGCM(testMsg, key, iv)
	if err != nil {
		t.Errorf("Runtime error on encrypt")
		return
	}
	decryptedVal, err := DecryptAESGCM(encryptedVal, key, iv)
	if err != nil {
		t.Errorf("Runtime error on decrypt")
		return
	}
	if decryptedVal != testMsg {
		t.Errorf("Expected decrypted to get '%s' but got '%s'", testMsg, decryptedVal)
		return
	}
}

func TestEncryptShouldGetTheCorrectResult(t *testing.T) {
	testMsg := "This value should be the same values + ภาษาไทย และตัวเลข 0123456789 ด้วย"
	key := []byte("59azblwa8d64f6a718p061156a21oi68")
	iv := make([]byte, 12)
	expectedBase64Encrypted := "hHd6PqJpVnUOdfzMQ8HuSwq1C90qPP+q2t2mIiqvoidWYvZjl96ow/xnhbxhS5bfRBGusiKCm9WJut6jecIU+vYfcMXxiDIUKQHlZRQW1/B+fXf2oEtvTlSvZC7m+EiupXBYSFqqfiW7ERNEftilBlK9Opa6y0jeWpS7ACHsO3U="
	encryptedVal, err := EncryptAESGCM(testMsg, key, iv)
	if err != nil {
		t.Errorf("Runtime error on encrypt")
		return
	}

	if encryptedVal != expectedBase64Encrypted {
		t.Errorf("Expected decrypted to get '%s' but got '%s'", expectedBase64Encrypted, encryptedVal)
		return
	}

}

func TestDecryptedShouldGetTheCorrectResult(t *testing.T) {
	key := []byte("59azblwa8d64f6a718p061156a21oi68")
	iv := make([]byte, 12)
	EncryptedString := "hHd6PqJpVnUOdfzMQ8HuSwq1C90qPP+q2t2mIiqvoidWYvZjl96ow/xnhbxhS5bfRBGusiKCm9WJut6jecIU+vYfcMXxiDIUKQHlZRQW1/B+fXf2oEtvTlSvZC7m+EiupXBYSFqqfiW7ERNEftilBlK9Opa6y0jeWpS7ACHsO3U="
	expectedDecrypted := "This value should be the same values + ภาษาไทย และตัวเลข 0123456789 ด้วย"
	decryptedVal, err := DecryptAESGCM(EncryptedString, key, iv)
	if err != nil {
		t.Errorf("Runtime error on encrypt")
		return
	}

	if decryptedVal != expectedDecrypted {
		t.Errorf("Expected decrypted to get '%s' but got '%s'", expectedDecrypted, decryptedVal)
		return
	}

}
