/*
Package ui implements the UI for the delorean application.
*/
package ui

import (
	"strings"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/ui/components/snapshots"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type components struct {
	tabs      *tabs.Model
	snapshots *snapshots.Model
}

type Model struct {
	components components
	snapshots  []domain.Snapshot
	keys       KeyMap
	state      *shared.State
	config     *config.Config
}

func NewModel(cfg *config.Config) (*Model, error) {
	st := shared.State{}
	tabsCmp, err := tabs.NewModel(&st, shared.GetTabItems())
	if err != nil {
		return &Model{}, err
	}
	snapshotsCmp, err := snapshots.NewModel()
	if err != nil {
		return &Model{}, err
	}

	snaps, err := btrfs.SnapshotsList(cfg.Volumes)
	if err != nil {
		return &Model{}, err
	}

	return &Model{
		components: components{
			tabs:      &tabsCmp,
			snapshots: &snapshotsCmp,
		},
		snapshots: snaps,
		keys:      getKeyMaps(),
		state:     &st,
		config:    cfg,
	}, nil
}

func (m *Model) Init() tea.Cmd {
	m.components.snapshots.UpdateList(m.snapshots)
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			cmd = tea.Quit
		}
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			m.onClick(msg)
		}
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	s := strings.Builder{}
	// tabs
	s.WriteString(m.components.tabs.View())
	s.WriteString("\n")

	// content
	if m.state.CurrentTab == shared.SnapshotsTab {
		mainContent := lipgloss.JoinHorizontal(lipgloss.Top, m.components.snapshots.View())
		s.WriteString(mainContent)
		s.WriteString("\n")
	} else if m.state.CurrentTab == shared.SettingsTab {
		s.WriteString(" settings\n")
	}

	return s.String()
}

func (m *Model) onClick(event tea.MouseMsg) {
	if event.Y <= tabs.TabsHeigh { // tabs area gets full length of the prompt
		m.components.tabs.OnClick(event)
	}
}
