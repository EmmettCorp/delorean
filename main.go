package main

import (
	"log"

	"github.com/EmmettCorp/delorean/pkg/app"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	a, err := app.New()
	if err != nil {
		return err
	}

	return a.Run()
}
