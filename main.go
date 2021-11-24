package main

import (
	"log"

	"github.com/EmmettCorp/delorean/pkg/app"
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/version"
)

const appName = "delorean"

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	appConfig, err := config.NewAppConfig(appName, version.Number)
	if err != nil {
		return err
	}

	app, err := app.NewApp(appConfig)
	if err != nil {
		return err
	}
	return app.Run()
}
