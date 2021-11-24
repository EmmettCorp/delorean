package app

import (
	"io"

	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/gui"
	"github.com/EmmettCorp/delorean/pkg/i18n"
	"github.com/EmmettCorp/delorean/pkg/log"
	"github.com/sirupsen/logrus"
)

// App struct
type App struct {
	closers []io.Closer

	Config       *config.AppConfig
	Log          *logrus.Entry
	OSCommand    *commands.OSCommand
	BtrfsCommand *commands.BtrfsCommand
	Gui          *gui.Gui
	Tr           *i18n.TranslationSet
	ErrorChan    chan error
}

// NewApp bootstrap a new application
func NewApp(config *config.AppConfig) (*App, error) {
	app := &App{
		closers:   []io.Closer{},
		Config:    config,
		ErrorChan: make(chan error),
	}
	var err error
	app.Log = log.NewLogger(config, "23432119147a4367abf7c0de2aa99a2d")
	app.Tr = i18n.NewTranslationSet(app.Log)
	app.OSCommand = commands.NewOSCommand(app.Log, config)

	// here is the place to make use of the docker-compose.yml file in the current directory

	app.BtrfsCommand, err = commands.NewBtrfsCommand(app.Log, app.OSCommand, app.Tr, app.Config, app.ErrorChan)
	if err != nil {
		return app, err
	}
	app.Gui, err = gui.New(app.Log)
	if err != nil {
		return app, err
	}
	return app, nil
}

func (app *App) Run() error {
	return app.Gui.Run()
}
