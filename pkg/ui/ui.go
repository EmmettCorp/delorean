/*
Package ui implements the UI for the delorean application.
*/
package ui

import (
	"strings"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/ui/components/snapshots"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type components struct {
	tabs      tea.Model
	snapshots tea.Model
}

type App struct {
	state      *shared.State
	components components
	keys       KeyMap
	config     *config.Config
}

func NewModel(cfg *config.Config) (*App, error) {
	st := shared.State{
		ActiveVolumes:     cfg.Volumes,
		ClickableElements: make(map[shared.TabItem][]shared.Clickable),
	}
	tabsCmp, err := tabs.NewModel(&st, shared.GetTabItems())
	if err != nil {
		return &App{}, err
	}
	snapshotsCmp, err := snapshots.NewModel(&st)
	if err != nil {
		return &App{}, err
	}

	a := App{
		components: components{
			tabs:      tabsCmp,
			snapshots: snapshotsCmp,
		},
		keys:   getKeyMaps(),
		config: cfg,
		state:  &st,
	}

	return &a, nil
}

func (a *App) Init() tea.Cmd {
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

	if a.state.CurrentTab == shared.SnapshotsTab {
		_, cmd = a.components.snapshots.Update(msg)
		cmds = append(cmds, cmd)
	}

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

func (a *App) OnClick(msg tea.MouseMsg) {
	clickable := a.state.FindClickable(msg.X, msg.Y)
	if clickable != nil {
		clickable.OnClick(msg)
	}
}
