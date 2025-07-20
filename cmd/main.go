package main

import (
	"govault/cli"
	"govault/internal/utils"
	"log"
)

func main() {
	utils.InitPPrint()
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
