package main

import (
	"govault/cli"
	"log"
)

func main() {
	// cobra cli init
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
