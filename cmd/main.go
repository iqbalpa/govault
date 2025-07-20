package main

import (
	"govault/cli"
	"log"
)

func main() {
	// // Init DB
	// db := utils.ConnectDb()
	// utils.MigrateDb(db)

	// masterPass := "iqbalpahlevi"
	// salt := crypto.GenerateRandomSalt()

	// sr := repository.New(db)
	// vs := service.New(*sr)
	// ar := repository.NewAuthRepo(db)
	// as := service.NewAuthService(*ar)

	// // 1. Vault init
	// as.InitMasterPass(masterPass)

	// // 2. Vault login
	// session, err := as.Login(masterPass)
	// if err != nil {
	// 	fmt.Println("master password is incorrect")
	// 	return
	// }
	// fmt.Println("session:\n", session)

	// // 3. Vault create
	// name := "gmail"
	// username := "iqbalpahlevi@gmail.com"
	// password := "malang123"
	// note := "ini gmail dummy aja"
	// s, _ := vs.CreateSecret(session.MasterPassword, name, username, password, note, salt)
	// fmt.Println("create secret:\n", s)

	// // 4. Vault get
	// id := s.ID
	// res, _ := vs.GetSecretById(session.MasterPassword, id)
	// fmt.Println("get by id:\n", res)

	// // 5. Vault list
	// res2, _ := vs.GetAllSecrets()
	// fmt.Println("get all:\n", res2)

	// // 6. Vault delete
	// vs.DeleteSecretById(id)

	// cobra cli init
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
