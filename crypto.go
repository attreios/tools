package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/osmontero/holmes"
)

func Encrypt(key string, plain string) string {
	h := holmes.New(os.Getenv("LOG_MODE"), "utils.Encrypt")

	plainBytes := []byte(plain)

	block, blockErr := aes.NewCipher([]byte(key))
	if blockErr != nil {
		h.FatalError("%v", blockErr)
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		h.FatalError("%v", gcmErr)
	}

	nonce := make([]byte, aesGCM.NonceSize())

	encrypted := aesGCM.Seal(nonce, nonce, plainBytes, nil)

	return fmt.Sprintf("%x", encrypted)
}

func GenerateCryptoKey(size int) string {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := fmt.Sprintf("%s", bytes)
	return key
}

func Decrypt(keySt string, encrypted string) string {
	h := holmes.New(os.Getenv("LOG_MODE"), "utils.Decrypt")

	key := []byte(keySt)
	enc, _ := hex.DecodeString(encrypted)

	block, ncErr := aes.NewCipher(key)
	if ncErr != nil {
		h.Error("%v", ncErr)
	}

	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		h.Error("%v", gcmErr)
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, openErr := aesGCM.Open(nil, nonce, ciphertext, nil)
	if openErr != nil {
		h.Error("%v", openErr)
	}

	return fmt.Sprintf("%s", plaintext)
}
