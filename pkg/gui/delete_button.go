package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/jroimartin/gocui"
)

func (gui *Gui) deleteButton() (*gocui.View, error) {
	buttonName := "delete"
	view, err := gui.g.SetView(gui.views.deleteBtn.name,
		gui.views.deleteBtn.x0,
		gui.views.deleteBtn.y0,
		gui.views.deleteBtn.x1,
		gui.views.deleteBtn.y1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.ErrLog.Printf("can't build %s button: %v", buttonName, err)

			return nil, err
		}
		err := gui.g.SetKeybinding(buttonName, gocui.MouseLeft, gocui.ModNone, gui.deleteSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, buttonName)
	}

	return view, nil
}

func (gui *Gui) deleteSnapshot(g *gocui.Gui, v *gocui.View) error {
	snap, err := gui.getChosenSnapshot()
	if err != nil {
		if errors.Is(err, domain.ErrSnapshotIsNotChosen) {
			gui.state.status = colors.FgRed(err.Error())

			return nil
		}

		return err
	}

	err = btrfs.DeleteSnapshot(snap.Path)
	if err != nil {
		return err
	}

	gui.state.status = colors.FgGreen(fmt.Sprintf("snapshot %s is deleted", snap.Label))
	err = gui.updateSnapshotsList()
	if err != nil {
		return err
	}

	return gui.escapeFromViewsByName(gui.views.snapshots.name)
}
