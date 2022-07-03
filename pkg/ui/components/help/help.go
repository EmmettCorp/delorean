package help

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	bbHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const fullHelpHeigh = 3

type Model struct {
	state *shared.State
	help  bbHelp.Model
	keys  shared.KeyMap
}

func New(state *shared.State) *Model {
	help := bbHelp.NewModel()
	help.Styles = bbHelp.Styles{
		ShortDesc:      helpTextStyle.Copy(),
		FullDesc:       helpTextStyle.Copy(),
		ShortSeparator: helpTextStyle.Copy(),
		FullSeparator:  helpTextStyle.Copy(),
		FullKey:        helpTextStyle.Copy(),
		ShortKey:       helpTextStyle.Copy(),
		Ellipsis:       helpTextStyle.Copy(),
	}

	return &Model{
		state: state,
		help:  help,
		keys:  shared.GetKeyMaps(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(msg, m.keys.Help) {
			m.help.ShowAll = !m.help.ShowAll
			if m.help.ShowAll {
				m.state.Areas.HelpBar.Height += fullHelpHeigh
			} else {
				m.state.Areas.HelpBar.Height -= fullHelpHeigh
			}
			m.state.ResizeAreas()
		}
	}

	return m, nil
}

func (m *Model) View() string {
	if !m.help.ShowAll {
		return helpStyle.Copy().
			Render(m.help.ShortHelpView(m.keys.ShortHelp()))
	}

	var kb [][]key.Binding
	if m.state.CurrentTab == shared.SnapshotsTab {
		kb = m.keys.SnapshotsHelp()
	} else if m.state.CurrentTab == shared.SettingsTab {
		kb = m.keys.SettingsHelp()
	}

	m.SetWidth(m.state.ScreenWidth)

	return helpStyle.Copy().
		Render(m.help.FullHelpView(kb))
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}
