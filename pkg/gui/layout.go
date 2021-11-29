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

	g.Highlight = true
	width, height := g.Size()

	minimumHeight := 9
	minimumWidth := 10

	var err error
	_, err = g.SetView("limit", 0, 0, width-1, height-1, 0)
	if err != nil && err.Error() != UNKNOWN_VIEW_ERROR_MSG {
		return err
	}
	gui.Views.Limit.Visible = height < minimumHeight || width < minimumWidth

	informationStr := gui.informationStr()
	appStatus := gui.statusManager.getStatusString()

	viewDimensions := gui.getWindowDimensions(informationStr, appStatus)

	// reading more lines into main view buffers upon resize
	prevMainView := gui.Views.Main
	if prevMainView != nil {
		_, prevMainHeight := prevMainView.Size()
		newMainHeight := viewDimensions["main"].Y1 - viewDimensions["main"].Y0 - 1
		heightDiff := newMainHeight - prevMainHeight
		if heightDiff > 0 {
			if manager, ok := gui.viewBufferManagerMap["main"]; ok {
				manager.ReadLines(heightDiff)
			}
			if manager, ok := gui.viewBufferManagerMap["secondary"]; ok {
				manager.ReadLines(heightDiff)
			}
		}
	}

	setViewFromDimensions := func(viewName string, windowName string, frame bool) (*gocui.View, error) {
		dimensionsObj, ok := viewDimensions[windowName]

		if !ok {
			// view not specified in dimensions object: so create the view and hide it
			// making the view take up the whole space in the background in case it needs
			// to render content as soon as it appears, because lazyloaded content (via a pty task)
			// cares about the size of the view.
			view, err := g.SetView(viewName, 0, 0, width, height, 0)
			if view != nil {
				view.Visible = false
			}
			return view, err
		}

		frameOffset := 1
		if frame {
			frameOffset = 0
		}
		view, err := g.SetView(
			viewName,
			dimensionsObj.X0-frameOffset,
			dimensionsObj.Y0-frameOffset,
			dimensionsObj.X1+frameOffset,
			dimensionsObj.Y1+frameOffset,
			0,
		)

		if view != nil {
			view.Visible = true
		}

		return view, err
	}

	return nil
}

func (gui *Gui) createAllViews() error {
	viewNameMappings := []struct {
		viewPtr **gocui.View
		name    string
	}{
		{viewPtr: &gui.Views.Main, name: "main"},
		{viewPtr: &gui.Views.Settings, name: "settings"},
		{viewPtr: &gui.Views.Prompt, name: "prompt"},
		{viewPtr: &gui.Views.Options, name: "options"},
		{viewPtr: &gui.Views.Limit, name: "limit"},
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

	gui.Views.Main.Highlight = true
	gui.Views.Main.FgColor = theme.GocuiDefaultTextColor

	return nil
}

func (gui *Gui) prepareView(viewName string) (*gocui.View, error) {
	// arbitrarily giving the view enough size so that we don't get an error, but
	// it's expected that the view will be given the correct size before being shown
	return gui.g.SetView(viewName, 0, 0, 10, 10, 0)
}
