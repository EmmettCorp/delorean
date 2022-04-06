/*
Package tabs keeps the logic for tabs component.
*/
package tabs

import (
	"errors"
	"os"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const (
	TabsHeigh            = 2
	tabsLeftRightIndents = 2
)

type Model struct {
	shared.Clickable

	state *shared.State
	tabs  []*Tab
}

func NewModel(state *shared.State, tabItems []shared.TabItem) (*Model, error) {
	if len(tabItems) == 0 {
		return nil, errors.New("empty tabItems")
	}

	m := Model{
		state: state,
	}

	m.X1, m.Y1, m.X2, m.Y2 = m.getCoords()

	var x1 int
	for i := range tabItems {
		title := tabItems[i].String()
		x2 := x1 + len(title) + 3 // nolint:gomnd // 3 = 2 vertical bars + 1 space
		nt := NewTab(state, tabItems[i], x1, 0, x2, TabsHeigh)
		m.tabs = append(m.tabs, nt)
		m.AddSuccessor(nt)
		x1 = x2 + 1
	}

	return &m, nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m *Model) GetCoords() (int, int, int, int) {
	return m.X1, m.Y1, m.X2, m.Y2
}

func (m *Model) getCoords() (int, int, int, int) {
	physicalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return -1, -1, -1, -1
	}

	return 0, 0, physicalWidth, TabsHeigh
}

func (m *Model) View() string {
	var tabs []string
	for i := range m.tabs {
		if m.state.CurrentTab == m.tabs[i].id {
			tabs = append(tabs, activeTab.Render(m.tabs[i].title))
		} else {
			tabs = append(tabs, inactiveTab.Render((m.tabs[i].title)))
		}
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tabs...,
	)

	physicalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err.Error()
	}
	gap := tabGap.Render(strings.Repeat(" ", max(0, physicalWidth-lipgloss.Width(row)-tabsLeftRightIndents)))

	return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
