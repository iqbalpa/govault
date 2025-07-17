package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Loading .env file...")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file")
		return
	}
	fmt.Println("Loaded .env successfully!")
}
