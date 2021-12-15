package gui

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) scheduleView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.schedule.name,
		int(0.8*float32(gui.maxX)),
		gui.views.schedule.y0,
		gui.maxX,
		gui.maxY-5,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.schedule.name, err)
			return nil, err
		}

		view.Title = gui.views.schedule.name
	}

	return view, nil
}
