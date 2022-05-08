/*
Package app is responsible for creating a new application instance.
*/
package app

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	App struct {
		gui *tea.Program
	}
)

// New creates and returns new app.
func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("can't get new config: %v", err)
	}

	model, err := ui.NewModel(cfg)
	if err != nil {
		return nil, err
	}
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	return &App{
		gui: p,
	}, nil
}

// Run setup and run gui handlers.
func (a *App) Run() error {
	return a.gui.Start()
}
