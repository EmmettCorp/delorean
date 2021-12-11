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

		views   views
		buttons buttons
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

	gui.initButtons()
	vv := gocui.ManagerFunc(gui.layout)
	// manager
	gui.g.SetManager(gui.buttons.create, gui.buttons.restore, gui.buttons.delete, gui.buttons.settings, vv)

	// keybindings
	bb := gui.GetInitialKeybindings()
	err := gui.setKeybindings(bb)
	if err != nil {
		return fmt.Errorf("can't set keybindings: %v", err)
	}

	return gui.g.MainLoop()
}

func (gui *Gui) Stop() {
	gui.g.Close()
}
