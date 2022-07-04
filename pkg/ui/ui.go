/*
Package ui implements the UI for the delorean application.
*/
package ui

import (
	"os"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/EmmettCorp/delorean/pkg/ui/components/help"
	"github.com/EmmettCorp/delorean/pkg/ui/components/settings"
	"github.com/EmmettCorp/delorean/pkg/ui/components/snapshots"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/alert"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const (
	minWidth = 81
	minHeigh = 28
)

type components struct {
	tabs      tea.Model
	snapshots tea.Model
	settings  tea.Model
	help      tea.Model
	alert     *alert.Model
}

type App struct {
	state      *shared.State
	components components
	keys       shared.KeyMap
	config     *config.Config
}

func NewModel(cfg *config.Config) (*App, error) {
	var err error
	st := shared.NewState(cfg)
	st.ScreenWidth, st.ScreenHeight, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return nil, err
	}

	tabsCmp, err := tabs.New(st, shared.GetTabItems())
	if err != nil {
		return &App{}, err
	}
	snapshotsCmp, err := snapshots.New(st)
	if err != nil {
		return &App{}, err
	}
	settingsCmp, err := settings.New(st)
	if err != nil {
		return &App{}, err
	}
	helpCmp := help.New(st)
	alertCmp := alert.New()

	a := App{
		components: components{
			tabs:      tabsCmp,
			snapshots: snapshotsCmp,
			settings:  settingsCmp,
			help:      helpCmp,
			alert:     alertCmp,
		},
		keys:   shared.GetKeyMaps(),
		config: cfg,
		state:  st,
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
		cmd = a.keyEventHandle(msg)
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			a.OnClick(msg)
		} else {
			a.components.snapshots.Update(msg)
		}
	case tea.WindowSizeMsg:
		a.onWindowSizeChanged(msg)
		a.components.snapshots.Update(msg)
		a.components.settings.Update(msg)
	}
	cmds = append(cmds, cmd)

	if a.state.UpdateSnapshots {
		a.components.snapshots.Update(msg)
	}

	return a, tea.Batch(cmds...)
}

func (a *App) View() string {
	s := strings.Builder{}

	if a.windowTooSmall() {
		a.components.alert.SetTitle("Window is too small.")

		return a.components.alert.View(a.state.ScreenWidth, a.state.ScreenHeight)
	}
	// tabs
	s.WriteString(a.components.tabs.View())
	s.WriteString("\n")

	// content
	if a.state.CurrentTab == shared.SnapshotsTab {
		mainContent := lipgloss.JoinHorizontal(lipgloss.Top, a.components.snapshots.View())
		s.WriteString(mainContent)
		s.WriteString("\n")
	} else if a.state.CurrentTab == shared.SettingsTab {
		s.WriteString(a.components.settings.View())
		s.WriteString("\n")
	}

	// help
	s.WriteString(a.components.help.View())

	return s.String()
}

func (a *App) keyEventHandle(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.Matches(msg, a.keys.Quit):
		return tea.Quit
	case key.Matches(msg, a.keys.Help):
		a.components.help.Update(msg)

		return nil
	case key.Matches(msg, a.keys.Tab, a.keys.ShiftTab):
		a.components.tabs.Update(msg)

		return nil
	default:
		return a.componentsKeyEventHandle(msg)
	}
}

func (a *App) componentsKeyEventHandle(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	switch a.state.CurrentTab {
	case shared.SnapshotsTab:
		_, cmd = a.components.snapshots.Update(msg)
	case shared.SettingsTab:
		_, cmd = a.components.settings.Update(msg)
	}

	return cmd
}

func (a *App) OnClick(msg tea.MouseMsg) {
	clickable := a.state.FindClickable(msg.X, msg.Y)
	if clickable != nil {
		// nolint:errcheck,gosec // TODO: handle error. Add to log + consider to write short message to status bar
		clickable.OnClick(msg)
	}
}

func (a *App) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	a.state.ScreenWidth = msg.Width
	a.state.ScreenHeight = msg.Height
	a.state.ResizeAreas()
	logger.Client.InfoLog.Println(a.state.ScreenWidth, a.state.ScreenHeight)
}

func (a *App) windowTooSmall() bool {
	if a.state.ScreenWidth < minWidth || a.state.ScreenHeight < minHeigh {
		return true
	}

	return false
}
