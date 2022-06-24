package settings

import (
	"fmt"
	"io"
	"strings"

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
func (v *volume) Height() int         { return 1 }
func (v *volume) Spacing() int        { return 1 }
func (v *volume) Update(msg tea.Msg, m *list.Model) tea.Cmd {
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

	fmt.Fprint(w, row.String())
}
