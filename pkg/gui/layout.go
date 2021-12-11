package gui

import "github.com/jroimartin/gocui"

// // layout is called for every screen re-render e.g. when the screen is resized
func (gui *Gui) layout(g *gocui.Gui) error {
	return gui.initViews()
}
