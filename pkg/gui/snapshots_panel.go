package gui

import (
	"errors"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) snapshotsView(maxX, maxY int) (*gocui.View, error) {
	name := "snapshots"
	view, err := gui.g.SetView(name, gui.indent, gui.headerHight, int(0.8*float32(maxX)), maxY-5)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", name, err)
			return nil, err
		}
	}

	view.Title = name

	return view, nil
}
