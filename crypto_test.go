package blind

import (
	"encoding/hex"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key, _ := GenerateCryptoKey(32)
	t.Log("key:", hex.EncodeToString(key))
	text := "Testing with some text"
	t.Log("to encrypt:", text)
	encrypted, _ := Encrypt(key, text)
	t.Log("encrypted:", encrypted)
	decrypted, _ := Decrypt(key, encrypted)
	t.Log("decrypted:", decrypted)
	if text != decrypted {
		t.Error("Text should be:", text)
	}
}
