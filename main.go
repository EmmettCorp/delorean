package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

const (
	indent      = 0
	headerHight = 2
)

// DummyWidget does nothing and is needed only to draw interface.
type DummyWidget struct{}

// ButtonWidget is a button widget.
type ButtonWidget struct {
	name    string
	x, y    int
	w       int
	label   string
	handler func(g *gocui.Gui, v *gocui.View) error
}

// Layout draws a new layout.
func (w *ButtonWidget) Layout(g *gocui.Gui) error {
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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("side-main", indent, headerHight, int(0.8*float32(maxX)), maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("main", int(0.8*float32(maxX)), headerHight, maxX, maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("terminal", indent, maxY-5, maxX, maxY); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	maxX, _ := g.Size()
	g.Mouse = true

	createBtn := NewButtonWidget("create", 0, -1, "create", DummyFunc(&DummyWidget{}))
	restoreBtn := NewButtonWidget("restore", 8, -1, "restore", DummyFunc(&DummyWidget{}))
	deleteBtn := NewButtonWidget("delete", 17, -1, "delete", DummyFunc(&DummyWidget{}))
	gearBtn := NewButtonWidget("settings", maxX-10, -1, "settings", DummyFunc(&DummyWidget{}))
	fl := gocui.ManagerFunc(layout)
	g.SetManager(createBtn, restoreBtn, deleteBtn, gearBtn, fl)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// NewButtonWidget creates a new button.
func NewButtonWidget(name string, x, y int, label string,
	handler func(g *gocui.Gui, v *gocui.View) error) *ButtonWidget {
	return &ButtonWidget{name: name, x: x, y: y, w: len(label), label: label, handler: handler}
}

// DummyFunc does nothing and is needed only to draw interface.
func DummyFunc(status *DummyWidget) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return nil
	}
}
