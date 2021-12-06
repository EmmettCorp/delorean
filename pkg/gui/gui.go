/*
Pakcage gui is responsible user interface.
*/
package gui

import "github.com/jroimartin/gocui"

// Gui wraps the gocui Gui object which handles rendering and events
type (
	Gui struct {
		g *gocui.Gui

		Views Views

		ViewsSetup bool
	}
)

// New creates and returns a new gui handler.
func New() (*Gui, error) {
	g := Gui{}
	return &g, nil
}
