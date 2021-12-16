package gui

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) snapshotsView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.snapshots.name,
		gui.views.snapshots.x0,
		gui.views.snapshots.y0,
		int(0.8*float32(gui.maxX)),
		gui.maxY-5,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.snapshots.name, err)
			return nil, err
		}

		view.Title = gui.views.snapshots.name
	}

	return view, nil
}
