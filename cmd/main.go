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

	sr := repository.New(db)
	vs := service.New(*sr)

	name := "gmail"
	username := "iqbalpahlevi@gmail.com"
	password := "malang123"
	note := "ini gmail dummy aja"

	s, _ := vs.CreateSecret(masterPass, name, username, password, note, salt)
	fmt.Println("create secret:\n", s)

	id := s.ID
	res, _ := vs.GetSecretById(masterPass, id)
	fmt.Println("get by id:\n", res)

	res2, _ := vs.GetAllSecrets()
	fmt.Println("get all:\n", res2)

	vs.DeleteSecretById(id)
}
