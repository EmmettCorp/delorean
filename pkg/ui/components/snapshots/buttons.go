package snapshots

import (
	"fmt"
	"path"
	"time"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/rate"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	createButtonHeight = 2
	createLimitInSec   = 2
)

type createButton struct {
	title  string
	coords shared.Coords
	state  *shared.State
	// Limiter for create button is needed to allow to finish create snapshot operation.
	// There is no real point in real life doing snapshots every second.
	// If allow user to call btrfs.CreateSnapshot several times a second it could cause a exec.Command call error.
	limiter *rate.Limiter
	// Update snapshots list calls btrfs call to get all the list of the chosen volumes. It's too expensive.
	// Instead we'll call it only when it's really needed. Ex.: create or delete item.
	updateCallback func()
}

func newCreateButton(st *shared.State, title string, coords shared.Coords, callback func()) *createButton {
	cb := createButton{
		state:          st,
		title:          title,
		limiter:        rate.NewLimiter(time.Second * createLimitInSec),
		updateCallback: callback,
	}

	cb.SetCoords(coords)

	return &cb
}

func (cb *createButton) SetTitle(title string) {
	cb.title = title
}

func (cb *createButton) GetTitle() string {
	return cb.title
}

func (cb *createButton) GetCoords() shared.Coords {
	return cb.coords
}

func (cb *createButton) SetCoords(c shared.Coords) {
	cb.coords = c
}

func (cb *createButton) OnClick(event tea.MouseMsg) error {
	if !cb.limiter.Allow() {
		// TODO: consider to return errors.New("too many create calls per second")
		// and write it to to status bar

		return nil
	}

	var activeVolumeFound bool

	for _, vol := range cb.state.Config.Volumes {
		if !vol.Active {
			continue
		}

		activeVolumeFound = true

		if cb.state.Config.VolumeInRootFs(vol) {
			err := btrfs.CreateSnapshot(vol.Device.MountPoint,
				path.Join(vol.SnapshotsPath, domain.Manual))
			if err != nil {
				return fmt.Errorf("can't create snapshot for %s: %v", vol.Device.MountPoint, err)
			}

			continue
		}

		err := btrfs.CreateSnapshot(vol.Device.MountPoint, path.Join(vol.SnapshotsPath, domain.Manual))
		if err != nil {
			return fmt.Errorf("can't create snapshot for %s: %v", vol.Device.MountPoint, err)
		}
	}

	if !activeVolumeFound {
		// TODO: after creation write message to the status bar
		// put the message to a status bar errors.New("there are no active volumes")

		return nil
	}

	cb.updateCallback()

	return nil
}
