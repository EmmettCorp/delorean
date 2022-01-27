package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/logger"
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

	err = gui.escapeFromEditableView(g, view)
	if err != nil {
		return err
	}
	gui.state.status = colors.FgGreen(fmt.Sprintf("%s data is saved ", view.Name()))

	return gui.updateSnapshotsList()
}

func (gui *Gui) escapeFromEditableView(g *gocui.Gui, view *gocui.View) error {
	view.Highlight = false
	view.SelBgColor = gocui.ColorDefault

	gui.setDefaultStatus()
	_, err := gui.g.SetCurrentView(gui.views.status.name)

	return err
}

func (gui *Gui) escapeFromView(g *gocui.Gui, view *gocui.View) error {
	gui.setDefaultStatus()
	_, err := gui.g.SetCurrentView(gui.views.status.name)

	return err
}

func (gui *Gui) escapeFromViewsByName(names ...string) error {
	for _, name := range names {
		view, err := gui.g.View(name)
		if err != nil {
			if !errors.Is(err, gocui.ErrUnknownView) {
				logger.Client.ErrLog.Printf("can't get %s view: %v", name, err)

				return err
			}
		}
		view.Highlight = false
		view.SelBgColor = gocui.ColorDefault
	}

	_, err := gui.g.SetCurrentView(gui.views.status.name)

	return err
}
