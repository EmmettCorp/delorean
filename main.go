package main

import (
	"log"

	"github.com/EmmettCorp/delorean/pkg/ui"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ui.Draw()
	return nil
	// closeLogger, err := logger.Init()
	// if err != nil {
	// 	return err
	// }

	// err = commands.CheckIfRoot()
	// if err != nil {
	// 	return err
	// }

	// a, err := app.New()
	// if err != nil {
	// 	return err
	// }

	// err = a.Run()
	// if err != nil {
	// 	logger.Client.ErrLog.Printf("main loop err: %v", err)
	// 	fmt.Printf("main loop err: %v", err) // nolint:forbidigo // on purpose here
	// }

	// return closeLogger()
}
