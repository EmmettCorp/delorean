package gui

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/logger"
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
			logger.Client.ErrLog.Printf("can't build %s button: %v", gui.views.restoreBtn.name, err)

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
			gui.state.status = colors.Paint(err.Error(), colors.Red)

			return nil
		}

		return err
	}

	vol, err := gui.getVolumeByID(snap.VolumeID)
	if err != nil {
		return err
	}

	if !gui.volumeInRootFs(vol) {
		gui.state.status = colors.Paint(
			fmt.Sprintf("volume %s is not a child of top level subvolume", vol.Label), colors.Red)

		return gui.escapeFromViewsByName(gui.views.snapshots.name)
	}

	err = btrfs.CreateSnapshot(vol.Device.MountPoint, path.Join(vol.SnapshotsPath, domain.Restore))
	if err != nil {
		return fmt.Errorf("can't create revert snapshot for %s: %v", vol.Device.MountPoint, err)
	}

	subvolumeDelorianMountPoint := path.Join(domain.DeloreanMountPoint, vol.Subvol)
	oldFsDelorianMountPoint := path.Join(domain.DeloreanMountPoint, fmt.Sprintf("%s.old", vol.Subvol))

	err = os.Rename(subvolumeDelorianMountPoint, oldFsDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't rename directory %s: %v", oldFsDelorianMountPoint, err)
	}

	err = btrfs.Restore(snap.Path, subvolumeDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't create snapshot for %s: %v", vol.Device.MountPoint, err)
	}

	gui.config.ToRemove = append(gui.config.ToRemove, oldFsDelorianMountPoint)
	if vol.Device.MountPoint == "/" {
		gui.config.Path = path.Join(subvolumeDelorianMountPoint, gui.config.Path)
	}
	err = gui.config.Save()
	if err != nil {
		return fmt.Errorf("can't add old path to remove: %v", err)
	}

	gui.state.status = colors.Paint("reboot system to compete restore", colors.Red)

	return gui.escapeFromViewsByName(gui.views.snapshots.name)
}
