package gui

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) storageView(maxX, maxY int) (*gocui.View, error) {
	name := "storage"
	view, err := gui.g.SetView(name, gui.indent, maxY-5, maxX, maxY)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", name, err)
			return nil, err
		}

		view.Title = name
	}

	return view, nil
}
