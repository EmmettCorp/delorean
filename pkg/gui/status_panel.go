package gui

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) statusView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.status.name,
		gui.views.status.x0,
		gui.views.status.y0,
		gui.maxX,
		gui.views.status.y1,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.status.name, err)
			return nil, err
		}

		view.Highlight = true
		view.SelFgColor = gocui.ColorGreen
	}

	view.Clear()
	fmt.Fprint(view, gui.state.status)

	return view, nil
}
