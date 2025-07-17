package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the .env file:", err)
	}
	fmt.Println("Sucessfully loaded .env")
}
