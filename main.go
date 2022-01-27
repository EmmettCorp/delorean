package main

import (
	"fmt"
	"log"

	"github.com/EmmettCorp/delorean/pkg/app"
	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/logger"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := logger.Init()
	if err != nil {
		return err
	}

	err = commands.CheckIfRoot()
	if err != nil {
		return err
	}

	a, err := app.New()
	if err != nil {
		return err
	}

	err = a.Run()
	if err != nil {
		logger.Client.ErrLog.Printf("main loop err: %v", err)
		fmt.Printf("main loop err: %v", err) // nolint forbidigo: on purpose here
	}

	return logger.Client.Close()
}
