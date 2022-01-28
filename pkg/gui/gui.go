/*
Package gui is responsible for user interface.
*/
package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/jroimartin/gocui"
)

// Gui wraps the gocui Gui object which handles rendering and events.
type (
	Gui struct {
		g *gocui.Gui

		views views

		config *config.Config

		state     *state
		maxX      int
		maxY      int
		snapshots []domain.Snapshot
	}
)

// New creates and returns a new gui handler.
func New(cfg *config.Config) (*Gui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		logger.Client.ErrLog.Printf("can't get new gui: %v", err)

		return nil, fmt.Errorf("can't get new gui: %v", err)
	}

	g.Mouse = true
	g.InputEsc = true

	return &Gui{
		g:      g,
		config: cfg,
		state:  initState(),
	}, nil
}

// Run setup the gui with keybindings and start the mainloop.
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
		logger.Client.ErrLog.Printf("can't set keybindings: %v", err)

		return fmt.Errorf("can't set keybindings: %v", err)
	}

	err = gui.g.MainLoop()
	if err != nil {
		if errors.Is(err, gocui.ErrQuit) {
			logger.Client.InfoLog.Print("quit")

			return nil
		}

		logger.Client.ErrLog.Printf("main loop failed: %v", err)
		for _, v := range gui.allViews() {
			err = quit(gui.g, v)
			if err != nil {
				return err
			}
		}
	}

	return err
}
