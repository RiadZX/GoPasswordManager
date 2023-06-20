package passwords

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func EncryptPassword(password, masterPassword string) (string, error) {
	plaintext := []byte(password)
	key := []byte(generateKey(masterPassword))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// Return the encrypted password as a base64-encoded string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptPassword(encryptedPassword, masterPassword string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return "", err
	}

	key := []byte(generateKey(masterPassword))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// Generate a 32-byte key using the master password, change this to be secure. this is unsafe now.
func generateKey(masterPassword string) string {
	key := []byte(masterPassword)
	for len(key) < 32 {
		key = append(key, key...)
	}
	return string(key[:32])
}
