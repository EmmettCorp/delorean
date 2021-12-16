package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) restoreButton() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.restoreBtn.name,
		gui.views.restoreBtn.x0,
		gui.views.restoreBtn.y0,
		gui.views.restoreBtn.x1,
		gui.views.restoreBtn.y1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", gui.views.restoreBtn.name, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(gui.views.restoreBtn.name, gocui.MouseLeft, gocui.ModNone, gui.restoreSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, gui.views.restoreBtn.name)
	}

	return view, nil
}

func (gui *Gui) restoreSnapshot(g *gocui.Gui, v *gocui.View) error {
	gui.state.status = " reboot system to compete resotre "
	return nil
}
