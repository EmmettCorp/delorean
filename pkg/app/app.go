/*
Package app is responsible for creating a new application instance.
*/
package app

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/bl"
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/storage"
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

	db, err := storage.New(cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("can't init storage: %v", err)
	}
	sr, err := storage.NewSnapshotRepo(db)
	if err != nil {
		return nil, fmt.Errorf("can't init snapshots repo: %v", err)
	}
	gr, err := storage.NewGarbageRepo(db)
	if err != nil {
		return nil, fmt.Errorf("can't init garbage repo: %v", err)
	}

	blService := bl.New(sr, gr, cfg)

	model, err := ui.NewModel(blService, cfg)
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
