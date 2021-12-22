package main

import (
	"log"

	"github.com/EmmettCorp/delorean/pkg/app"
	"github.com/EmmettCorp/delorean/pkg/commands"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := commands.CheckIfRoot()
	if err != nil {
		return err
	}

	app, err := app.New()
	if err != nil {
		return err
	}

	return app.Run()
}
