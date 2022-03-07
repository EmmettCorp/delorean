package gui

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/jroimartin/gocui"
)

const (
	helpViewWidth = 50
	helpViewHeigh = 20
)

func (gui *Gui) helpView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.helpView.name,
		gui.maxX/2-helpViewWidth/2,
		gui.maxX/2+helpViewWidth/2,
		gui.maxY/2-helpViewHeigh/2,
		gui.maxY/2+helpViewHeigh/2+helpViewHeigh%2,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't set %s view: %v", gui.views.errorView.name, err)

			return nil, err
		}

		view.Editable = false
		view.Frame = true
		view.Wrap = false
		gui.g.Cursor = false
		view.Highlight = true

		out := &bytes.Buffer{}
		for _, kb := range gui.GetInitialKeybindings() {
			fmt.Fprintf(out, fmt.Sprintf("%s %s\n",
				kb.Keys, fmt.Sprintf(fmt.Sprintf("%%%ds", helpViewWidth-len(kb.Keys)-4), kb.Description)))
		}
	}

	return view, nil
}

func (gui *Gui) deleteHelpView() error {
	_, err := gui.g.View(gui.views.errorView.name)
	if err != nil {
		return nil // nolint
	}

	return gui.g.DeleteView(gui.views.errorView.name)
}
