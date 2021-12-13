package gui

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) statusView(maxX, maxY int) (*gocui.View, error) {
	name := "status"
	view, err := gui.g.SetView(name, gui.buttons.width, -1, maxX, gui.headerHight-1)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", name, err)
			return nil, err
		}

		view.Highlight = true
		view.SelFgColor = gocui.ColorGreen
	}

	view.Clear()
	fmt.Fprint(view, gui.state.status)

	return view, nil
}
