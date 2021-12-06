package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) createAllViews() error {
	viewNameMappings := []struct {
		viewPtr **gocui.View
		name    string
	}{
		{viewPtr: &gui.Views.Main, name: "main"},
		{viewPtr: &gui.Views.Secondary, name: "secondary"},
	}

	var err error
	for _, mapping := range viewNameMappings {
		*mapping.viewPtr, err = gui.g.SetView(mapping.name, 0, 0, 10, 10)
		if err != nil && err.Error() != errUnknownView {
			return err
		}
	}

	return nil
}

// layout is called for every screen re-render e.g. when the screen is resized
func (gui *Gui) layout(g *gocui.Gui) error {
	if !gui.ViewsSetup {
		if err := gui.createAllViews(); err != nil {
			return err
		}
	}

	g.Highlight = true
	width, height := g.Size()

	minimumHeight := 9
	minimumWidth := 10

	var err error
	_, err = g.SetView("limit", 0, 0, width-1, height-1)
	if err != nil && err.Error() != errUnknownView {
		return err
	}

	fmt.Println(width, height, minimumWidth, minimumHeight)
	return nil
}
