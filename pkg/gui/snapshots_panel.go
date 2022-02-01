package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/jroimartin/gocui"
)

type snapshotEditor struct {
	g *Gui
}

func (gui *Gui) snapshotsView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.snapshots.name,
		gui.views.snapshots.x0,
		gui.views.snapshots.y0,
		int(scheduleIndent*float32(gui.maxX)),
		gui.maxY-volumesHeigh,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't set %s view: %v", gui.views.snapshots.name, err)

			return nil, err
		}

		view.Title = gui.views.snapshots.name
		view.Editor = &snapshotEditor{
			g: gui,
		}

		err := gui.setSnapshotsKeybindings()
		if err != nil {
			return nil, err
		}

		err = gui.updateSnapshotsList()
		if err != nil {
			return nil, err
		}
	}

	return view, nil
}

func (gui *Gui) setSnapshotsKeybindings() error {
	err := gui.g.SetKeybinding(gui.views.snapshots.name, gocui.MouseLeft, gocui.ModNone, gui.editSnapshots)
	if err != nil {
		return err
	}
	err = gui.g.SetKeybinding(gui.views.snapshots.name, gocui.KeyEsc, gocui.ModNone, gui.escapeFromEditableView)
	if err != nil {
		return err
	}
	err = gui.g.SetKeybinding(gui.views.snapshots.name, gocui.KeyDelete, gocui.ModNone, gui.deleteSnapshot)
	if err != nil {
		return err
	}

	return nil
}

func (gui *Gui) updateSnapshotsList() error {
	view, err := gui.g.View(gui.views.snapshots.name)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't get %s view: %v", gui.views.snapshots.name, err)

			return err
		}
	}
	view.Clear()
	snaps, err := btrfs.SnapshotsList(gui.config.Volumes)
	if err != nil {
		return err
	}
	for i := range snaps {
		fmt.Fprintf(view, " %s\t%s\t%s \n", snaps[i].Label, snaps[i].Type, snaps[i].VolumeLabel)
	}
	gui.snapshots = snaps

	return nil
}

func (gui *Gui) getChosenSnapshot() (domain.Snapshot, error) {
	view, err := gui.g.View(gui.views.snapshots.name)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't get %s view: %v", gui.views.snapshots.name, err)

			return domain.Snapshot{}, err
		}
	}

	v := gui.g.CurrentView()
	if v == nil || v.Name() != view.Name() {
		return domain.Snapshot{}, domain.ErrSnapshotIsNotChosen
	}

	_, cY := view.Cursor()

	return gui.snapshots[cY], nil
}

func (gui *Gui) editSnapshots(g *gocui.Gui, view *gocui.View) error {
	err := gui.escapeFromViewsByName(gui.views.schedule.name)
	if err != nil {
		return err
	}

	gui.g.Cursor = false
	view.Highlight = true
	view.Editable = true
	view.SelBgColor = gocui.ColorBlack

	_, err = gui.g.SetCurrentView(gui.views.snapshots.name)
	if err != nil {
		return err
	}

	return err
}

func (e *snapshotEditor) Edit(view *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	if key == gocui.KeyArrowDown {
		_, cY := view.Cursor()
		if cY >= len(e.g.snapshots)-1 {
			return
		}
		view.MoveCursor(0, 1, false)
	}

	if key == gocui.KeyArrowUp {
		_, cY := view.Cursor()
		if cY <= 0 {
			return
		}
		view.MoveCursor(0, -1, false)
	}
}
