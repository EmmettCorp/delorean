package gui

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
)

const (
	active   = "✔"
	inactive = "✗"
)

func (gui *Gui) volumesView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.volumes.name,
		gui.views.volumes.x0,
		gui.maxY-5,
		gui.maxX,
		gui.maxY,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.volumes.name, err)
			return nil, err
		}

		view.Title = gui.views.volumes.name
	}
	view.Clear()
	gui.drawVolumes(view)

	return view, nil
}

func (gui *Gui) drawVolumes(view *gocui.View) {
	for i := range gui.config.Volumes {
		fmt.Fprintf(view, "Label: %s UID: %s ", gui.config.Volumes[i].Label, gui.config.Volumes[i].UID)
		if gui.config.Volumes[i].Active {

			fmt.Fprint(view, color.GreenString("%s", active))
		} else {
			fmt.Fprint(view, color.RedString("%s", inactive))
		}
	}
}
