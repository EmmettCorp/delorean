package gui

import (
	"github.com/jroimartin/gocui"
)

type Binding struct {
	Name        string
	ViewName    string
	Contexts    []string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         gocui.Key
	Modifier    gocui.Modifier
	Description string
	Alternative string
	Tag         string // e.g. 'navigation'. Used for grouping things in the cheatsheet
	OpensMenu   bool
}

// GetInitialKeybindings is a function.
func (gui *Gui) GetInitialKeybindings() []*Binding {
	bindings := []*Binding{
		{
			Name:        "Ctrl+h",
			ViewName:    "",
			Key:         gocui.KeyCtrlH,
			Modifier:    gocui.ModNone,
			Handler:     gui.toggleHelp,
			Description: "Show/hide this help dialog",
		},
		{
			Name:        "Ctrl+c",
			ViewName:    "",
			Key:         gocui.KeyCtrlC,
			Modifier:    gocui.ModNone,
			Handler:     quit,
			Description: "Quit delorean",
		},
		{
			Name:        "Ctrl+q",
			ViewName:    "",
			Key:         gocui.KeyCtrlQ,
			Modifier:    gocui.ModNone,
			Handler:     quit,
			Description: "Quit delorean",
		},
		{
			Name:        "Esc",
			ViewName:    "",
			Key:         gocui.KeyEsc,
			Modifier:    gocui.ModNone,
			Handler:     gui.escapeFromView,
			Description: "Move focus from current view",
		},
		{
			Name:     "",
			ViewName: "",
			Key:      gocui.KeyDelete,
			Modifier: gocui.ModNone,
			Handler:  dummyHandler,
		},
	}

	return bindings
}

func (gui *Gui) setKeybindings(bindings []*Binding) error {
	for b := range bindings {
		err := gui.g.SetKeybinding(
			bindings[b].ViewName,
			bindings[b].Key,
			bindings[b].Modifier,
			bindings[b].Handler,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func dummyHandler(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (gui *Gui) toggleHelp(g *gocui.Gui, v *gocui.View) error {
	if gui.views.errorView.visible {
		return nil
	}

	if gui.views.helpView.visible {
		return gui.deleteHelpView()
	}

	_, err := gui.helpView()

	return err
}
