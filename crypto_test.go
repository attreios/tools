package tools

import (
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key, _ := GenerateCryptoKey(32) 
	text := "Testing with some text"
	encrypted, _ := Encrypt(key, text)
	decrypted, _ := Decrypt(key, encrypted)

	if text != decrypted {
		t.Error("Text should be:", text)
	}
}
