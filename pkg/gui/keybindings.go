package gui

import (
	"github.com/jroimartin/gocui"
)

type Binding struct {
	Keys        string
	ViewName    string
	Contexts    []string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         gocui.Key
	Modifier    gocui.Modifier
	Description []byte
	Alternative string
	Tag         string // e.g. 'navigation'. Used for grouping things in the cheatsheet
	OpensMenu   bool
}

// GetInitialKeybindings is a function.
func (gui *Gui) GetInitialKeybindings() []*Binding {
	bindings := []*Binding{
		{
			Keys:     "Ctrl+C",
			ViewName: "",
			Key:      gocui.KeyCtrlC,
			Modifier: gocui.ModNone,
			Handler:  quit,
		},
		{
			Keys:     "Ctrl+Q",
			ViewName: "",
			Key:      gocui.KeyCtrlQ,
			Modifier: gocui.ModNone,
			Handler:  quit,
		},
		{
			Keys:     "Esc",
			ViewName: "",
			Key:      gocui.KeyEsc,
			Modifier: gocui.ModNone,
			Handler:  gui.escapeFromView,
		},
		{
			Keys:     "Del",
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
