/*
Package settings is responsible for settings logic.
*/
package settings

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	state  *shared.State
	height int
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

	res := docStyle.Height(m.height).Render(content)

	return res[:len(res)-len(content)] // will be changes in future
}

func (m *Model) getHeight() int {
	return m.state.Areas.MainContent.Height
}
