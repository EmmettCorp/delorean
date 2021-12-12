package gui

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) scheduleView(maxX, maxY int) (*gocui.View, error) {
	name := "schedule"
	view, err := gui.g.SetView(name, int(0.8*float32(maxX)), gui.headerHight, maxX, maxY-5)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", name, err)
			return nil, err
		}
	}

	view.Title = name

	return view, nil
}
