package gui

import (
	"github.com/EmmettCorp/delorean/pkg/theme"
	"github.com/jesseduffield/gocui"
)

// layout is called for every screen re-render e.g. when the screen is resized
func (gui *Gui) layout(g *gocui.Gui) error {
	// the main logic to draw interface here.
	if !gui.ViewsSetup {
		if err := gui.createAllViews(); err != nil {
			return err
		}
	}
	return nil
}

func (gui *Gui) createAllViews() error {
	viewNameMappings := []struct {
		viewPtr **gocui.View
		name    string
	}{
		{viewPtr: &gui.Views.Main, name: "main"},
		{viewPtr: &gui.Views.Main, name: "settings"},
		{viewPtr: &gui.Views.Main, name: "prompt"},
		// {viewPtr: &gui.Views.Options, name: "options"},
		// {viewPtr: &gui.Views.Credentials, name: "credentials"},
		// {viewPtr: &gui.Views.Menu, name: "menu"},
		// {viewPtr: &gui.Views.Confirmation, name: "confirmation"},
	}

	var err error
	for _, mapping := range viewNameMappings {
		*mapping.viewPtr, err = gui.prepareView(mapping.name)
		if err != nil && err.Error() != UNKNOWN_VIEW_ERROR_MSG {
			return err
		}
	}

	gui.Views.Options.Frame = false
	gui.Views.Options.FgColor = theme.OptionsColor

	return nil
}

func (gui *Gui) prepareView(viewName string) (*gocui.View, error) {
	// arbitrarily giving the view enough size so that we don't get an error, but
	// it's expected that the view will be given the correct size before being shown
	return gui.g.SetView(viewName, 0, 0, 10, 10, 0)
}
