package config

import "github.com/jroimartin/gocui"

type Colors struct {
	Background gocui.Attribute
	Foreground gocui.Attribute
	Highlight  gocui.Attribute
}
