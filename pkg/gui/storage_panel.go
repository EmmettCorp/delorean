package gui

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) storageView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.storage.name,
		gui.views.storage.x0,
		gui.maxY-5,
		gui.maxX,
		gui.maxY,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.storage.name, err)
			return nil, err
		}

		view.Title = gui.views.storage.name
	}

	return view, nil
}
