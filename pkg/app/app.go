/*
Package app is responsible for creating a new application instance.
*/
package app

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/gui"
	"github.com/EmmettCorp/delorean/pkg/logger"
)

type (
	App struct {
		gui *gui.Gui
	}
)

// New creates and returns new app.
func New() (*App, error) {
	log, err := logger.New()
	if err != nil {
		return nil, fmt.Errorf("can't get new logger: %v", err)
	}

	cfg, err := config.New(log)
	if err != nil {
		return nil, fmt.Errorf("can't get new config: %v", err)
	}

	g, err := gui.New(cfg, log)
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
