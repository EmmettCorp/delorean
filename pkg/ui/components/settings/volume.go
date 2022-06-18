package settings

import (
	"fmt"
	"io"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	active           = "✔"
	inactive         = "✘"
	volumeNameColumn = "Volumes"
	volumeInfoColumn = "Info"
	minColumnGap     = "  "
)

var defStyles = list.NewDefaultItemStyles()

type (
	volume struct {
		ID               string
		Subvol           string
		Label            string
		SnapshotsPath    string
		Active           bool
		DeviceUUID       string
		DevicePath       string
		DeviceMountPoint string
	}
)

// FilterValue is used to set filter item and required for `list.Model` interface.
func (v *volume) FilterValue() string { return v.Label }
func (d *volume) Height() int         { return 1 }
func (d *volume) Spacing() int        { return 1 }
func (d *volume) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (v *volume) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	s, ok := listItem.(*volume)
	if !ok {
		return
	}

	sign := signInactiveStyle.Render(inactive)
	if s.Active {
		sign = signActiveStyle.Render(active)
	}
	var row strings.Builder
	if index == m.Index() {
		row.WriteString(fmt.Sprintf("%s %s", selectedItem.Render(s.Label), sign))
	} else {
		row.WriteString(fmt.Sprintf("%s %s", normalItem.Render(s.Label), sign))
	}

	fmt.Fprintf(w, row.String())
}

func domainVolumeToVolume(dv domain.Volume) volume {
	return volume{
		ID:               dv.ID,
		Subvol:           dv.Subvol,
		Label:            dv.Label,
		SnapshotsPath:    dv.SnapshotsPath,
		Active:           dv.Active,
		DeviceUUID:       dv.Device.UUID,
		DevicePath:       dv.Device.Path,
		DeviceMountPoint: dv.Device.MountPoint,
	}
}
