package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) restoreButton() (*gocui.View, error) {
	view, err := gui.g.SetView(restoreButton,
		gui.buttons.restoreButton.x0,
		-1,
		gui.buttons.restoreButton.x1,
		gui.headerHight-1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", restoreButton, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(restoreButton, gocui.MouseLeft, gocui.ModNone, gui.restoreSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, restoreButton)
	}

	return view, nil
}

func (gui *Gui) restoreSnapshot(g *gocui.Gui, v *gocui.View) error {
	gui.state.status = " reboot system to compete resotre "
	return nil
}
