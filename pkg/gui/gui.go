/*
Package gui provides helpers to work with gui.
*/
package gui

import (
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/utils"
	"github.com/jesseduffield/gocui"
	"github.com/sirupsen/logrus"
)

// Gui wraps the gocui Gui object which handles rendering.
type Gui struct {
	g      *gocui.Gui
	Log    *logrus.Entry
	Config *config.AppConfig
}

// New creates a new gui handler.
func New(log *logrus.Entry) (*Gui, error) {
	return &Gui{
		Log: log,
	}, nil
}

func (gui Gui) Run() error {
	g := gocui.NewGui()

	defer g.Close()

	gui.g = g

	if err := gui.SetColorScheme(); err != nil {
		return err
	}
	return nil
}

// SetColorScheme sets the color scheme for the app based on the user config
func (gui *Gui) SetColorScheme() error {
	gui.g.FgColor = gui.GetColor(gui.Config.UserConfig.Gui.Theme.InactiveBorderColor)
	gui.g.SelFgColor = gui.GetColor(gui.Config.UserConfig.Gui.Theme.ActiveBorderColor)
	return nil
}

// GetColor bitwise OR's a list of attributes obtained via the given keys
func (gui *Gui) GetColor(keys []string) gocui.Attribute {
	var attribute gocui.Attribute
	for _, key := range keys {
		attribute |= utils.GetGocuiAttribute(key)
	}
	return attribute
}
