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
		err = gui.g.SetKeybinding(gui.views.volumes.name, gocui.MouseLeft, gocui.ModNone, gui.editVolumes)
		if err != nil {
			return nil, err
		}
		err = gui.g.SetKeybinding(gui.views.volumes.name, gocui.KeyEnter, gocui.ModNone, gui.saveConfig)
		if err != nil {
			return nil, err
		}

		gui.drawVolumes(view)
	}

	return view, nil
}

func (gui *Gui) drawVolumes(view *gocui.View) {
	view.Clear()
	for i := range gui.config.Volumes {
		if !gui.config.Volumes[i].Mounted {
			continue
		}
		var activeSign string
		if gui.config.Volumes[i].Active {
			activeSign = colors.FgGreen(active)
		} else {
			activeSign = colors.FgRed(inactive)
		}
		fmt.Fprintf(view, " %s %s\n",
			gui.config.Volumes[i].Label, activeSign)
	}
}

func (gui *Gui) editVolumes(g *gocui.Gui, view *gocui.View) error {
	err := gui.escapeFromViewsByName(gui.views.schedule.name, gui.views.snapshots.name)
	if err != nil {
		return err
	}

	view.Editable = true
	_, cY := view.Cursor()
	if cY < len(gui.config.Volumes) {
		gui.config.Volumes[cY].Active = !gui.config.Volumes[cY].Active
		gui.drawVolumes(view)
		gui.state.status = colors.FgRed("press enter to save volumes settings")
	}

	_, err = gui.g.SetCurrentView(gui.views.volumes.name)
	return err
}
