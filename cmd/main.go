package main

import (
	"fmt"
	"govault/internal/crypto"
	"govault/internal/repository"
	"govault/internal/service"
	"govault/internal/utils"
)

func main() {
	db := utils.ConnectDb()
	utils.MigrateDb(db)

	masterPass := "iqbalpahlevi"
	salt := crypto.GenerateRandomSalt()
	derivedKey, _ := crypto.DeriveKey(masterPass, salt)

	sr := repository.New(db)
	vs := service.New(*sr)

	name := "gmail"
	username := "iqbalpahlevi@gmail.com"
	password := "malang123"
	note := "ini gmail dummy aja"

	vs.CreateSecret(name, username, password, note, derivedKey)

	id := "734cb362-4340-4e77-8118-978cb6995c87"
	res, _ := vs.GetSecretById(id)
	fmt.Println(res)

	res2, _ := vs.GetAllSecrets()
	fmt.Println(res2)
}
