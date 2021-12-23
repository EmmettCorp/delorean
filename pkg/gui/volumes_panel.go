package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/colors"
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
		var activeSign string
		if gui.config.Volumes[i].Active {
			activeSign = colors.FgGreen(active)
		} else {
			activeSign = colors.FgRed(inactive)
		}
		fmt.Fprintf(view, "Label: %s UID: %s %s\n",
			gui.config.Volumes[i].Label, gui.config.Volumes[i].UID, activeSign)
	}
}
