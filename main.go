package main

import (
	"log"

	"github.com/EmmettCorp/delorean/pkg/app"
	"github.com/EmmettCorp/delorean/pkg/logger"
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

	closeLogger, err := logger.Init()
	if err != nil {
		return err
	}
	defer func() {
		logErr := closeLogger()
		if err == nil && logErr != nil {
			err = logErr
		}
	}()

	err = a.Run()

	return err
}
