package utils

import (
	"fmt"
	"govault/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// default value
var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func init() {
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
}

func ConnectDb() *gorm.DB {
	fmt.Println("Connecting to DB...")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to db")
		return nil
	}
	fmt.Println("Connected to DB successfully!")
	return db
}

func MigrateDb(db *gorm.DB) {
	fmt.Println("Running DB migrations...")
	db.AutoMigrate(&model.Secret{})
	db.AutoMigrate(&model.Auth{})
	fmt.Println("DB Migrations is succeed!")
}
