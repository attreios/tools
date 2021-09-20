package blind

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Encrypt(key string, plain string) (string, error) {
	plainBytes := []byte(plain)

	block, blockErr := aes.NewCipher([]byte(key))
	if blockErr != nil {
		return "", blockErr
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		return "", gcmErr
	}

	nonce := make([]byte, aesGCM.NonceSize())

	encrypted := aesGCM.Seal(nonce, nonce, plainBytes, nil)

	return fmt.Sprintf("%x", encrypted), nil
}

func GenerateCryptoKey(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}	
	return string(bytes), nil
}

func Decrypt(keySt string, encrypted string) (string, error) {
	key := []byte(keySt)
	enc, _ := hex.DecodeString(encrypted)

	block, ncErr := aes.NewCipher(key)
	if ncErr != nil {
		return "", ncErr
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		return "", gcmErr
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, openErr := aesGCM.Open(nil, nonce, ciphertext, nil)
	if openErr != nil {
		return "", openErr
	}

	return string(plaintext), nil
}
