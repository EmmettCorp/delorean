package main

import (
	"log"

	"github.com/EmmettCorp/delorean/pkg/app"
)

const appName = "delorean"

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app, err := app.New()
	if err != nil {
		return err
	}

	return app.Run()
}
