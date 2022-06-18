/*
Package settings is responsible for settings logic.
*/
package settings

import (
	"fmt"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/divider"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	state        *shared.State
	list         list.Model
	coords       shared.Coords
	keys         keyMap
	longestLabel int
	err          error
}

func New(st *shared.State) (*Model, error) {
	m := Model{
		state: st,
		keys:  getKeyMaps(),
		coords: shared.Coords{
			X1: 0,
			Y1: st.Areas.TabBar.Height + 1,
			X2: st.Areas.MainContent.Width,
			Y2: st.Areas.MainContent.Height,
		},
	}

	itemsModel := list.New([]list.Item{}, &volume{}, 0, 0)
	itemsModel.SetFilteringEnabled(false)
	itemsModel.SetShowFilter(false)
	itemsModel.SetShowTitle(false)
	itemsModel.SetShowStatusBar(false)
	itemsModel.SetShowHelp(false)
	m.list = itemsModel
	m.UpdateList()

	return &m, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Enter) {
			m.toggleActive()
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	var s strings.Builder
	s.WriteString("\n")
	s.WriteString(styles.MainDocStyle.Foreground(styles.DefaultTheme.InactiveText).Render(m.getHeader()))
	s.WriteString("\n")
	s.WriteString(divider.HorizontalLine(m.state.ScreenWidth, styles.DefaultTheme.InactiveText))
	s.WriteString("\n")
	m.list.SetSize(10, m.state.Areas.MainContent.Height-4)

	s.WriteString(
		styles.MainDocStyle.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				listStyle.Render(m.list.View()),
				m.renderMetadata(),
			),
		),
	)
	s.WriteString("\n")
	s.WriteString(divider.HorizontalLine(m.state.ScreenWidth, styles.DefaultTheme.InactiveText))
	return s.String()
}

func (m *Model) toggleActive() {
	item := m.list.SelectedItem()
	s, ok := item.(*volume)
	if !ok {
		return
	}

	s.Active = !s.Active
	// TODO: save to config
}

func (m *Model) renderMetadata() string {
	item := m.list.SelectedItem()
	s, ok := item.(*volume)
	if !ok {
		return ""
	}

	status := "not active: snapshots for this subvolume are hidden"
	if s.Active {
		status = "active: snapshots for this subvolume are shown"
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		fmt.Sprintf("Subvolume label:\t\t%s\n", s.Label),
		fmt.Sprintf("Subvolume ID:\t\t%s\n", s.ID),
		fmt.Sprintf("Subvolume name:\t\t%s\n", s.Subvol),
		fmt.Sprintf("Device mount point:\t%s\n", s.DeviceMountPoint),
		fmt.Sprintf("Device UUID:\t\t%s\n", s.DeviceUUID),
		fmt.Sprintf("Device path:\t\t%s\n", s.DevicePath),
		fmt.Sprintf("Snapshots path:\t\t%s\n", s.SnapshotsPath),
		fmt.Sprintf("Status:\t\t\t%s", status),
	)
}

func (m *Model) UpdateList() {
	m.longestLabel = 0
	vols, err := m.getVolumes()
	if err != nil {
		m.err = err

		return
	}

	items := make([]list.Item, len(vols))
	for i := range vols {
		items[i] = &volume{
			ID:               vols[i].ID,
			Subvol:           vols[i].Subvol,
			Label:            vols[i].Label,
			SnapshotsPath:    vols[i].SnapshotsPath,
			Active:           vols[i].Active,
			DeviceUUID:       vols[i].Device.UUID,
			DevicePath:       vols[i].Device.Path,
			DeviceMountPoint: vols[i].Device.MountPoint,
		}
		if len(vols[i].Label) > m.longestLabel {
			m.longestLabel = len(vols[i].Label)
		}
	}

	m.list.SetItems(items)
}

func (m *Model) getVolumes() ([]domain.Volume, error) {
	// get volumes from database
	vv := []domain.Volume{ // dummy data
		{
			ID:            "540",
			Subvol:        "@",
			Label:         "Root",
			SnapshotsPath: "/run/delorean/.snapshots/@",
			Active:        true,
			Device: domain.Device{
				UUID:       "e81c9e92-2bba-4fa1-9a30-7f950532b051",
				Path:       "/dev/nvme0n1p2",
				MountPoint: "/",
			},
		},
		{
			ID:            "541",
			Subvol:        "@/home",
			Label:         "Home",
			SnapshotsPath: "/run/delorean/.snapshots/@home",
			Active:        false,
			Device: domain.Device{
				UUID:       "t29384d-2bba-as0d9f-08sf-134j434j324",
				Path:       "/dev/nvme0n1p3",
				MountPoint: "/home",
			},
		},
	}

	return vv, nil
}

func (m *Model) getHeader() string {
	var header strings.Builder
	header.WriteString(minColumnGap)
	header.WriteString(volumeNameColumn)
	header.WriteString(strings.Repeat(" ", m.longestLabel))
	header.WriteString(volumeInfoColumn)

	return header.String()
}
