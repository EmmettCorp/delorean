package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/version"
	"github.com/jroimartin/gocui"
)

func (gui *Gui) statusView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.status.name,
		gui.views.status.x0,
		gui.views.status.y0,
		gui.maxX,
		gui.views.status.y1,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.ErrLog.Printf("can't set %s view: %v", gui.views.status.name, err)

			return nil, err
		}
	}

	view.Clear()
	fmt.Fprintf(view, " %s", gui.state.status)

	return view, nil
}

func (gui *Gui) setDefaultStatus() {
	gui.state.status = fmt.Sprintf("delorean version %s | type ctrl+h to call help", version.Number)
}
