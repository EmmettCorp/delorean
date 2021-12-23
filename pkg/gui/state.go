package gui

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/version"
	"github.com/jroimartin/gocui"
)

type state struct {
	status string
}

func initState() *state {
	return &state{
		status: fmt.Sprintf(" delorean version %s | type ctrl+h to call help", version.Number),
	}
}

func (gui *Gui) saveConfig(g *gocui.Gui, view *gocui.View) error {
	err := gui.config.Save()
	if err != nil {
		return fmt.Errorf("can't save config: %v", err)
	}

	gui.escapeFromEditableView(g, view)
	gui.state.status = fmt.Sprintf(" %s data is saved ", view.Name())
	return nil
}

func (gui *Gui) escapeFromEditableView(g *gocui.Gui, view *gocui.View) error {
	view.Highlight = false
	view.SelBgColor = gocui.ColorDefault

	gui.setDefaultStatus()
	gui.g.SetCurrentView(gui.views.status.name)
	return nil
}
