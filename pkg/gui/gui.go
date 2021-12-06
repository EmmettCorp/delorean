/*
Pakcage gui is responsible user interface.
*/
package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// Gui wraps the gocui Gui object which handles rendering and events
type (
	Gui struct {
		g *gocui.Gui

		Views Views

		// this tells us whether our views have been initially set up
		ViewsSetup bool
	}
)

// New creates and returns a new gui handler.
func New() (*Gui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, fmt.Errorf("can't get new gui: %v", err)
	}

	return &Gui{
		g: g,
	}, nil
}

// Run setup the gui with keybindings and start the mainloop
func (gui *Gui) Run() error {
	// close gocui.Gui on close
	defer gui.g.Close()

	gui.g.SetManager(gocui.ManagerFunc(gui.layout))

	return gui.g.MainLoop()
}

func (gui *Gui) Stop() {
	gui.g.Close()
}
