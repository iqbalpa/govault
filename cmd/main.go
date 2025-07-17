package main

import (
	"fmt"
	"govault/internal/crypto"
	_ "govault/internal/crypto"
)

func main() {
	fmt.Println("hi")
	a, _ := crypto.EncryptAES("iqbalpahlevi")
	fmt.Println(a)
	b, _ := crypto.DecryptAES(a)
	fmt.Println(b)
}
