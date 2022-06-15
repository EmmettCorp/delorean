/*
Package settings is responsible for settings logic.
*/
package settings

import (
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	state   *shared.State
	height  int
	volumes list.Model
}

func NewModel(st *shared.State) (*Model, error) {
	m := Model{
		state: st,
	}

	m.height = m.getHeight()

	return &m, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(tea.WindowSizeMsg); ok {
		m.height = m.getHeight()
	}

	return m, nil
}

func (m *Model) View() string {
	content := " settings" // Just dummy "settings" string for now.

	res := styles.MainDocStyle.Height(m.height).Render(content)

	return res[:len(res)-len(content)] // will be changes in future
}

func (m *Model) getHeight() int {
	return m.state.Areas.MainContent.Height
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
