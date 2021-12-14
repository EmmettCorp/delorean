package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) createButton() (*gocui.View, error) {
	view, err := gui.g.SetView(createButton,
		gui.buttons.createButton.x0,
		-1,
		gui.buttons.createButton.x1,
		gui.headerHight-1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", createButton, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(createButton, gocui.MouseLeft, gocui.ModNone, gui.createSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, createButton)
	}

	return view, nil
}

func (gui *Gui) createSnapshot(g *gocui.Gui, v *gocui.View) error {
	gui.state.status = " new snapshot is created "
	return nil
}
