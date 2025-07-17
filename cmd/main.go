package main

import (
	"fmt"
	"govault/internal/crypto"
)

func main() {
	masterPass := "iqbal"
	salt := crypto.GenerateRandomSalt()
	key, _ := crypto.DeriveKey(masterPass, salt)
	a, _ := crypto.EncryptAES("iqbalpahlevi", key)
	fmt.Println(a)
	b, _ := crypto.DecryptAES(a, key)
	fmt.Println(b)
	c, _ := crypto.DeriveKey("iqbalpahlevi", salt)
	fmt.Println(c)
}
