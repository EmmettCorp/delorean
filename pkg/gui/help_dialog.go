package gui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/jroimartin/gocui"
)

const (
	helpViewWidth = 55
	helpViewHeigh = 20
	contentLength = helpViewWidth - 5
	two           = 2
)

func (gui *Gui) helpView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.helpView.name,
		gui.maxX/two-helpViewWidth/two,
		gui.maxY/two-helpViewHeigh/two,
		gui.maxX/two+helpViewWidth/two,
		gui.maxY/two+helpViewHeigh/two+helpViewHeigh%two,
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
		divider := strings.Repeat("-", helpViewWidth-two)

		fmt.Fprint(view, " General keybindings\n")
		fmt.Fprintf(view, "%s\n", divider)
		for _, kb := range gui.GetInitialKeybindings() {
			if kb.Name == "" {
				continue
			}
			fmt.Fprint(view, getKeybindingDescription(kb))
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

func getKeybindingDescription(kb *Binding) string {
	return fmt.Sprintf(" %s %s\n",
		colors.Paint(kb.Name, colors.Yellow),
		fmt.Sprintf(
			fmt.Sprintf("%%%ds", contentLength-len(kb.Name)), kb.Description),
	)
}
