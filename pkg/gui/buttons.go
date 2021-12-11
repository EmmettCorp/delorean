package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type buttons struct {
	create   *buttonWidget
	restore  *buttonWidget
	delete   *buttonWidget
	settings *buttonWidget
	width    int
}

type dummyWidget struct{}

type buttonWidget struct {
	name    string
	x, y    int
	w       int
	label   string
	handler func(g *gocui.Gui, v *gocui.View) error
}

// Layout draws a new layout.
func (w *buttonWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w+1, w.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView(w.name); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, w.handler); err != nil {
			return err
		}
		fmt.Fprint(v, w.label)
	}

	return nil
}

// NewButtonWidget creates a new button.
func NewButtonWidget(name string, x, y int, label string,
	handler func(g *gocui.Gui, v *gocui.View) error) *buttonWidget {
	return &buttonWidget{name: name, x: x, y: y, w: len(label), label: label, handler: handler}
}

// DummyFunc does nothing and is needed only to draw interface.
func DummyFunc(status *dummyWidget) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return nil
	}
}

func (gui *Gui) initButtons() {
	indent := 2
	create := "create"
	gui.buttons.create = NewButtonWidget(create, gui.buttons.width, -1, create, DummyFunc(&dummyWidget{}))
	gui.buttons.width += len(create) + indent
	restore := "restore"
	gui.buttons.restore = NewButtonWidget(restore, gui.buttons.width, -1, restore, DummyFunc(&dummyWidget{}))
	gui.buttons.width += len(restore) + indent
	delete := "delete"
	gui.buttons.delete = NewButtonWidget(delete, gui.buttons.width, -1, delete, DummyFunc(&dummyWidget{}))
	gui.buttons.width += len(delete) + indent
	settings := "settings"
	gui.buttons.settings = NewButtonWidget(settings, gui.buttons.width, -1, settings, DummyFunc(&dummyWidget{}))
	gui.buttons.width += len(settings) + indent
}
