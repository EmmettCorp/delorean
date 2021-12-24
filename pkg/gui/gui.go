/*
Pakcage gui is responsible user interface.
*/
package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/jroimartin/gocui"
	"go.uber.org/zap"
)

// Gui wraps the gocui Gui object which handles rendering and events
type (
	Gui struct {
		g *gocui.Gui

		views views

		config *config.Config

		log *zap.SugaredLogger

		state       *state
		headerHight int
		maxX        int
		maxY        int
		snapshots   []domain.Snapshot
	}
)

// New creates and returns a new gui handler.
func New(config *config.Config, log *zap.SugaredLogger) (*Gui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Errorf("can't get new gui: %v", err)
		return nil, fmt.Errorf("can't get new gui: %v", err)
	}

	g.Mouse = true
	g.InputEsc = true

	return &Gui{
		g:           g,
		config:      config,
		log:         log,
		state:       initState(),
		headerHight: 2,
	}, nil
}

// Run setup the gui with keybindings and start the mainloop
func (gui *Gui) Run() error {
	// close gocui.Gui on exit from main loop.
	defer gui.g.Close()

	gui.initViews()

	// manager
	gui.g.SetManager(gocui.ManagerFunc(gui.layout))

	// keybindings
	bb := gui.GetInitialKeybindings()
	err := gui.setKeybindings(bb)
	if err != nil {
		gui.log.Errorf("can't set keybindings: %v", err)
		return fmt.Errorf("can't set keybindings: %v", err)
	}

	err = gui.g.MainLoop()
	if err != nil {
		if errors.Is(err, gocui.ErrQuit) {
			gui.log.Info("quit")
			return nil
		}

		gui.log.Errorf("main loop failed: %v", err)
		for _, v := range gui.allViews() {
			quit(gui.g, v)
		}
	}

	return err
}
