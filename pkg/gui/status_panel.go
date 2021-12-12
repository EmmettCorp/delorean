package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) statusView(maxX, maxY int) (*gocui.View, error) {
	view, err := gui.g.SetView("status", gui.buttons.width, -1, maxX, headerHight-1)
	if err != nil && err != gocui.ErrUnknownView {
		gui.log.Errorf("can't set status view: %v", err)
		return nil, err
	}

	view.Clear()
	view.Highlight = true
	view.SelFgColor = gocui.ColorGreen
	fmt.Fprint(view, gui.state.status)

	return view, nil
}
