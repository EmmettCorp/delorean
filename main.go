package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(
			stop,
			syscall.SIGHUP,  // kill -SIGHUP XXXX
			syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
			syscall.SIGQUIT, // kill -SIGQUIT XXXX
		)
		<-stop
		app.Stop()
	}()

	return app.Run()
}
