package main

import (
	"fmt"
	"log"

	"github.com/EmmettCorp/delorean/pkg/config"
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
		tea.WithMouseCellMotion(),
	)
	if err := p.Start(); err != nil {
		return err
	}

	return nil
}
