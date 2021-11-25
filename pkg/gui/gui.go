/*
Package gui provides helpers to work with gui.
*/
package gui

import (
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/i18n"
	"github.com/EmmettCorp/delorean/pkg/utils"
	"github.com/jesseduffield/gocui"
	"github.com/sirupsen/logrus"
)

// OverlappingEdges determines if panel edges overlap
var OverlappingEdges = false

// Gui wraps the gocui Gui object which handles rendering.
type Gui struct {
	g      *gocui.Gui
	Log    *logrus.Entry
	Config *config.AppConfig
	Tr     *i18n.TranslationSet
	Views  Views
}

type Views struct {
	Main         *gocui.View
	Secondary    *gocui.View
	Options      *gocui.View
	Confirmation *gocui.View
	Menu         *gocui.View
	Credentials  *gocui.View
	Extras       *gocui.View
}

// A Manager is in charge of GUI's layout and can be used to build widgets.
type Manager interface {
	// Layout is called every time the GUI is redrawn, it must contain the
	// base views and its initializations.
	Layout(*Gui) error
}

// New creates a new gui handler.
func New(log *logrus.Entry) (*Gui, error) {
	return &Gui{
		Log: log,
	}, nil
}

var RuneReplacements = map[rune]string{}

func (gui Gui) Run() error {
	playMode := gocui.NORMAL

	g, err := gocui.NewGui(gocui.OutputTrue, OverlappingEdges, playMode, headless(), RuneReplacements)
	if err != nil {
		return err
	}

	defer g.Close()

	gui.g = g
	g.Mouse = true

	if err := gui.SetColorScheme(); err != nil {
		return err
	}

	g.SetManager(gocui.ManagerFunc(gui.layout), gocui.ManagerFunc(gui.getFocusLayout()))

	gui.Log.Info("starting main loop")

	return g.MainLoop()
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
