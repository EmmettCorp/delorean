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

type App struct {
	shared.Clickable

	state      *shared.State
	components components
	snapshots  []domain.Snapshot
	keys       KeyMap
	config     *config.Config
}

func NewModel(cfg *config.Config) (*App, error) {
	st := shared.State{}
	tabsCmp, err := tabs.NewModel(&st, shared.GetTabItems())
	if err != nil {
		return &App{}, err
	}
	snapshotsCmp, err := snapshots.NewModel()
	if err != nil {
		return &App{}, err
	}

	snaps, err := btrfs.SnapshotsList(cfg.Volumes)
	if err != nil {
		return &App{}, err
	}

	a := App{
		components: components{
			tabs:      tabsCmp,
			snapshots: &snapshotsCmp,
		},
		snapshots: snaps,
		keys:      getKeyMaps(),
		config:    cfg,
		state:     &st,
	}

	err = a.AddSuccessor(tabsCmp)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) Init() tea.Cmd {
	a.components.snapshots.UpdateList(a.snapshots)
	return nil
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, a.keys.Quit) {
			cmd = tea.Quit
		}
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			a.OnClick(msg)
		}
	}
	cmds = append(cmds, cmd)

	_, cmd = a.components.snapshots.Update(msg)
	cmds = append(cmds, cmd)

	return a, tea.Batch(cmds...)
}

func (a *App) View() string {
	s := strings.Builder{}
	// tabs
	s.WriteString(a.components.tabs.View())
	s.WriteString("\n")

	// content
	if a.state.CurrentTab == shared.SnapshotsTab {
		mainContent := lipgloss.JoinHorizontal(lipgloss.Top, a.components.snapshots.View())
		s.WriteString(mainContent)
		s.WriteString("\n")
	} else if a.state.CurrentTab == shared.SettingsTab {
		s.WriteString(" settings\n")
	}

	return s.String()
}
