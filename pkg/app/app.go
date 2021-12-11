package app

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/gui"
)

type (
	App struct {
		gui *gui.Gui
	}
)

// New creates and returns new app.
func New() (*App, error) {
	g, err := gui.New()
	if err != nil {
		return nil, fmt.Errorf("can't get new gui: %v", err)
	}
	return &App{
		gui: g,
	}, nil
}

// Run setup and run gui handlers.
func (a *App) Run() error {
	return a.gui.Run()
}

func (a *App) Stop() {
	a.gui.Stop()
}
