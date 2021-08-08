package utils

import (
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key, text := GenerateCryptoKey(32), "Testing with some text"
	encrypted := Encrypt(key, text)
	decrypted := Decrypt(key, encrypted)

	if text != decrypted {
		t.Error("Text should be:", text)
	}
}
