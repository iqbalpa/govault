package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	_ "govault/internal/utils"

	"golang.org/x/crypto/scrypt"
)

func createAesgcm(keyStr string) cipher.AEAD {
	key, err := base64.StdEncoding.DecodeString(keyStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	return aesgcm
}

// EncryptAES encrypts a string using AES-GCM.
func EncryptAES(text, key string) ([]byte, error) {
	nonce := make([]byte, 12)
	rand.Read(nonce)

	aesgcm := createAesgcm(key)

	plaintext := []byte(text)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	final := append(nonce, ciphertext...)
	return final, nil
}

// DecryptAES decrypts a string encrypted with AES-GCM.
func DecryptAES(ciphertext []byte, key string) (string, error) {
	nonce := ciphertext[:12]
	cipher := ciphertext[12:]

	aesgcm := createAesgcm(key)

	plaintext, err := aesgcm.Open(nil, nonce, cipher, nil)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(plaintext), nil
}

// DeriveKey derives a key from a master password using scrypt.
func DeriveKey(masterPass string, salt []byte) (string, error) {
	key, err := scrypt.Key([]byte(masterPass), salt, 32768, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

// GenerateRandomSalt generates a random salt for key derivation.
func GenerateRandomSalt() []byte {
	salt := make([]byte, 16)
	rand.Read(salt)
	return salt
}
