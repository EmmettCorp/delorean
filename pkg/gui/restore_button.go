package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/jroimartin/gocui"
)

func (gui *Gui) restoreButton() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.restoreBtn.name,
		gui.views.restoreBtn.x0,
		gui.views.restoreBtn.y0,
		gui.views.restoreBtn.x1,
		gui.views.restoreBtn.y1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", gui.views.restoreBtn.name, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(gui.views.restoreBtn.name, gocui.MouseLeft, gocui.ModNone, gui.restoreSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, gui.views.restoreBtn.name)
	}

	return view, nil
}

func (gui *Gui) restoreSnapshot(g *gocui.Gui, v *gocui.View) error {
	snap, err := gui.getChosenSnapshot()
	if err != nil {
		if errors.Is(err, domain.ErrSnapshotIsNotChosen) {
			gui.state.status = colors.FgRed(err.Error())
			return nil
		}
		return err
	}

	err = commands.SetDefault(snap.VolumePoint, snap.ID)
	if err != nil {
		return err
	}

	gui.state.status = colors.FgRed("reboot system to compete restore")

	return gui.escapeFromViewsByName(gui.views.snapshots.name)
}
