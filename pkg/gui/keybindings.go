package gui

import (
	"github.com/jroimartin/gocui"
)

type Binding struct {
	ViewName    string
	Contexts    []string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         interface{} // FIXME: find out how to get `gocui.Key | rune`
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
			ViewName: "",
			Key:      gocui.KeyCtrlC,
			Modifier: gocui.ModNone,
			Handler:  quit,
		},
		{
			ViewName: "",
			Key:      gocui.KeyCtrlQ,
			Modifier: gocui.ModNone,
			Handler:  quit,
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
