package blind

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

func Encrypt(key []byte, plain string) (string, error) {
	plainBytes := []byte(plain)

	block, blockErr := aes.NewCipher(key)
	if blockErr != nil {
		return "", blockErr
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		return "", gcmErr
	}

	nonce := make([]byte, aesGCM.NonceSize())

	encrypted := aesGCM.Seal(nonce, nonce, plainBytes, nil)

	return hex.EncodeToString(encrypted), nil
}

func GenerateCryptoKey(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func Decrypt(key []byte, encrypted string) (string, error) {
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

	decrypted, openErr := aesGCM.Open(nil, nonce, ciphertext, nil)
	if openErr != nil {
		return "", openErr
	}

	return string(decrypted), nil
}
