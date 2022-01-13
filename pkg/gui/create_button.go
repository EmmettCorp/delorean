package gui

import (
	"fmt"
	"path"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/jroimartin/gocui"
)

func (gui *Gui) createButton() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.createBtn.name,
		gui.views.createBtn.x0,
		gui.views.createBtn.y0,
		gui.views.createBtn.x1,
		gui.views.createBtn.y1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", gui.views.createBtn.name, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(gui.views.createBtn.name, gocui.MouseLeft, gocui.ModNone, gui.createSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, gui.views.createBtn.name)
	}

	return view, nil
}

func (gui *Gui) createSnapshot(g *gocui.Gui, view *gocui.View) error {
	if !gui.views.createBtn.limiter.Allow() {
		gui.state.status = colors.FgRed("too many create calls per second")
		return nil
	}

	var activeVolumeFound bool

	for _, vol := range gui.config.Volumes {
		if !vol.Active || !vol.Mounted {
			continue
		}

		activeVolumeFound = true

		err := commands.CreateSnapshot(vol.MountPoint, path.Join(vol.SnapshotsPath, domain.Manual))
		if err != nil {
			return fmt.Errorf("can't create snapshot for %s: %v", vol.MountPoint, err)
		}
	}

	if !activeVolumeFound {
		gui.state.status = colors.FgRed("there are no active volumes")
		return nil
	}

	gui.state.status = colors.FgGreen("new snapshot is created")

	return gui.updateSnapshotsList()
}
