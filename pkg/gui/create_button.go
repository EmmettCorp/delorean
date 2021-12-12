package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) createButton() (*gocui.View, error) {
	buttonName := "create"
	x := gui.buttons.width
	view, err := gui.g.SetView(buttonName, x, -1, x+len(buttonName)+1, gui.headerHight-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", buttonName, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(buttonName, gocui.MouseLeft, gocui.ModNone, gui.createSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, buttonName)
	}

	gui.buttons.width += len(buttonName) + gui.buttons.indent

	return view, nil
}

func (gui *Gui) createSnapshot(g *gocui.Gui, v *gocui.View) error {
	gui.state.status = " new snapshot is created "
	return nil
}
