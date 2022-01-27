/*
Package app is responsible for creating a new application instance.
*/
package app

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/gui"
)

type (
	App struct {
		gui *gui.Gui
	}
)

// New creates and returns new app.
func New() (*App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("can't get new config: %v", err)
	}

	g, err := gui.New(cfg)
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
