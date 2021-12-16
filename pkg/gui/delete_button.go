package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) deleteButton() (*gocui.View, error) {
	buttonName := "delete"
	view, err := gui.g.SetView(gui.views.deleteBtn.name,
		gui.views.deleteBtn.x0,
		gui.views.deleteBtn.y0,
		gui.views.deleteBtn.x1,
		gui.views.deleteBtn.y1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", buttonName, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(buttonName, gocui.MouseLeft, gocui.ModNone, gui.deleteSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, buttonName)
	}

	return view, nil
}

func (gui *Gui) deleteSnapshot(g *gocui.Gui, v *gocui.View) error {
	gui.state.status = " snapshot is deleted "
	return nil
}
