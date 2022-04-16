/*
Package ui implements the UI for the delorean application.
*/
package ui

import (
	"os"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/ui/components/help"
	"github.com/EmmettCorp/delorean/pkg/ui/components/snapshots"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type components struct {
	tabs      tea.Model
	snapshots tea.Model
	help      tea.Model
}

type App struct {
	state      *shared.State
	components components
	keys       shared.KeyMap
	config     *config.Config
}

func NewModel(cfg *config.Config) (*App, error) {
	var err error
	st := shared.State{
		ClickableElements: make(map[shared.TabItem][]shared.Clickable),
		Config:            cfg,
	}
	st.ScreenWidth, st.ScreenHeight, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	tabsCmp, err := tabs.NewModel(&st, shared.GetTabItems())
	if err != nil {
		return &App{}, err
	}
	snapshotsCmp, err := snapshots.NewModel(&st)
	if err != nil {
		return &App{}, err
	}
	helpCmp := help.NewModel(&st)

	a := App{
		components: components{
			tabs:      tabsCmp,
			snapshots: snapshotsCmp,
			help:      helpCmp,
		},
		keys:   shared.GetKeyMaps(),
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
	case tea.WindowSizeMsg:
		a.onWindowSizeChanged(msg)
	}
	cmds = append(cmds, cmd)

	// // tabs
	a.components.tabs.Update(msg)

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
	s.WriteString(a.components.help.View())

	return s.String()
}

func (a *App) OnClick(msg tea.MouseMsg) {
	clickable := a.state.FindClickable(msg.X, msg.Y)
	if clickable != nil {
		clickable.OnClick(msg)
	}
}

func (a *App) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	a.state.ScreenWidth = msg.Width
	a.state.ScreenHeight = msg.Height
}
