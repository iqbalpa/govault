package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	_ "govault/internal/utils"
)

var secretKey string
var key []byte
var aesgcm cipher.AEAD

func init() {
	secretKey = os.Getenv("SECRET_KEY")
	key = []byte(secretKey)
	if len(key) != 32 {
		panic("Key length must be 32")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err = cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
}

func EncryptAES(text string) (string, error) {
	nonce := make([]byte, 12)
	rand.Read(nonce)

	plaintext := []byte(text)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	final := append(nonce, ciphertext...)
	res := base64.StdEncoding.EncodeToString(final)
	return res, nil
}

func DecryptAES(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		fmt.Println(err.Error())
	}

	nonce := data[:12]
	cipher := data[12:]

	plaintext, err := aesgcm.Open(nil, nonce, cipher, nil)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(plaintext), nil
}

func DeriveKey(password string, salt int) (string, error) {
	return "", nil
}

func GenerateRandomSalt() int {
	return 1
}
