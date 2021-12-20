package gui

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) errorView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.errorView.name,
		gui.views.errorView.x0,
		gui.views.errorView.y0,
		gui.maxX-1,
		gui.maxY-1,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.errorView.name, err)
			return nil, err
		}

		view.Frame = true
		view.Wrap = false
		gui.g.Cursor = false
		view.Highlight = true
		view.SelFgColor = gocui.ColorRed
		fmt.Fprintln(view, "Terminal window is too small")
	}

	return view, nil
}

func (gui *Gui) deleteErrorView() error {
	_, err := gui.g.View(gui.views.errorView.name)
	if err != nil {
		return nil
	}

	return gui.g.DeleteView(gui.views.errorView.name)
}
