package snapshots

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	CreateButtonHeight = 2
)

type createButton struct {
	title  string
	coords shared.Coords
	state  *shared.State
	// err    error
}

func newCreateButton(st *shared.State, title string, coords shared.Coords) *createButton {
	cb := createButton{
		state: st,
		title: title,
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

func (cb *createButton) OnClick(event tea.MouseMsg) {
	cb.state.TestBool = !cb.state.TestBool
	// for _, vol := range cb.state.Config.Volumes {
	// 	if !vol.Active {
	// 		continue
	// 	}

	// 	if cb.state.Config.VolumeInRootFs(vol) {
	// 		err := btrfs.CreateSnapshot(vol.Device.MountPoint,
	// 			path.Join(vol.SnapshotsPath, domain.Manual))
	// 		if err != nil {
	// 			cb.err = fmt.Errorf("can't create snapshot for %s: %v", vol.Device.MountPoint, err)
	// 		}

	// 		continue
	// 	}

	// 	err := btrfs.CreateSnapshot(vol.Device.MountPoint, path.Join(vol.SnapshotsPath, domain.Manual))
	// 	if err != nil {
	// 		cb.err = fmt.Errorf("can't create snapshot for %s: %v", vol.Device.MountPoint, err)
	// 	}
	// }
}
