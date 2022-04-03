package main

import (
	"fmt"
	"log"

	"github.com/EmmettCorp/delorean/pkg/app"
	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/EmmettCorp/delorean/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("can't get new config: %v", err)
	}
	model, err := ui.NewModel(cfg)
	if err != nil {
		return err
	}
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}

func runOldUI() error {
	closeLogger, err := logger.Init()
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
		fmt.Printf("main loop err: %v", err) // nolint:forbidigo // on purpose here
	}

	return closeLogger()
}
