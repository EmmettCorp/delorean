package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/version"
	"github.com/jroimartin/gocui"
)

type state struct {
	status string
}

func initState() *state {
	return &state{
		status: fmt.Sprintf("delorean version %s | type ctrl+h to call help", version.Number),
	}
}

func (gui *Gui) saveConfig(g *gocui.Gui, view *gocui.View) error {
	err := gui.config.Save()
	if err != nil {
		return fmt.Errorf("can't save config: %v", err)
	}

	gui.escapeFromEditableView(g, view)
	gui.state.status = colors.FgGreen(fmt.Sprintf("%s data is saved ", view.Name()))
	return gui.updateSnapshotsList()
}

func (gui *Gui) escapeFromEditableView(g *gocui.Gui, view *gocui.View) error {
	view.Highlight = false
	view.SelBgColor = gocui.ColorDefault

	gui.setDefaultStatus()
	gui.g.SetCurrentView(gui.views.status.name)
	return nil
}

func (gui *Gui) escapeFromView(g *gocui.Gui, view *gocui.View) error {
	gui.setDefaultStatus()
	gui.g.SetCurrentView(gui.views.status.name)
	return nil
}

func (gui *Gui) escapeFromViewsByName(names ...string) error {
	for _, name := range names {
		view, err := gui.g.View(name)
		if err != nil {
			if !errors.Is(err, gocui.ErrUnknownView) {
				gui.log.Errorf("can't get %s view: %v", name, err)
				return err
			}
		}
		view.Highlight = false
		view.SelBgColor = gocui.ColorDefault
	}

	gui.g.SetCurrentView(gui.views.status.name)
	return nil
}
