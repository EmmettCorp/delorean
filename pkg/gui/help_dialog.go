package gui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/jroimartin/gocui"
)

const (
	helpViewWidth = 55
	helpViewHeigh = 20
)

func (gui *Gui) helpView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.helpView.name,
		gui.maxX/2-helpViewWidth/2,
		gui.maxY/2-helpViewHeigh/2,
		gui.maxX/2+helpViewWidth/2,
		gui.maxY/2+helpViewHeigh/2+helpViewHeigh%2,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't set %s view: %v", gui.views.helpView.name, err)

			return nil, err
		}

		view.Editable = false
		view.Frame = true
		view.Wrap = false
		gui.g.Cursor = false
		view.Highlight = true
		divider := strings.Repeat("-", helpViewWidth-2)

		fmt.Fprint(view, " General keybindings\n")
		fmt.Fprintf(view, "%s\n", divider)
		for _, kb := range gui.GetInitialKeybindings() {
			if kb.Name == "" {
				continue
			}
			fmt.Fprintf(view, fmt.Sprintf(" %s %s\n",
				kb.Name, fmt.Sprintf(fmt.Sprintf("%%%ds", helpViewWidth-len(kb.Name)-5), kb.Description)))
		}
	}
	gui.views.helpView.visible = true

	return view, nil
}

func (gui *Gui) deleteHelpView() error {
	_, err := gui.g.View(gui.views.helpView.name)
	if err != nil {
		return nil // nolint
	}
	gui.views.helpView.visible = false

	return gui.g.DeleteView(gui.views.helpView.name)
}
