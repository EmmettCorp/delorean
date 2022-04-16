package help

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	bbHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	state *shared.State
	help  bbHelp.Model
	keys  shared.KeyMap
}

func NewModel(state *shared.State) *Model {
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, nil
}

func (m *Model) View() string {
	return helpStyle.Copy().
		Width(m.state.ScreenWidth).
		Render(m.help.View(m.keys))
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}
