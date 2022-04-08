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

type clickableTab interface {
	shared.Clickable
	getTitle() string
	getID() shared.TabItem
}

type Model struct {
	state *shared.State
	Tabs  []clickableTab
}

func NewModel(state *shared.State, tabItems []shared.TabItem) (*Model, error) {
	if len(tabItems) == 0 {
		return nil, errors.New("empty tabItems")
	}

	m := Model{
		state: state,
	}

	var x1 int
	for i := range tabItems {
		title := tabItems[i].String()
		x2 := x1 + len(title) + 3 // nolint:gomnd // 3 = 2 vertical bars + 1 space
		nt := NewTab(state, tabItems[i], shared.Coords{
			X1: x1,
			Y1: 0,
			X2: x2,
			Y2: TabsHeigh,
		})
		m.Tabs = append(m.Tabs, nt)
		x1 = x2 + 1
	}

	return &m, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Model) View() string {
	var tabs []string
	for i := range m.Tabs {
		if m.state.CurrentTab == m.Tabs[i].getID() {
			tabs = append(tabs, activeTab.Render(m.Tabs[i].getTitle()))
		} else {
			tabs = append(tabs, inactiveTab.Render((m.Tabs[i].getTitle())))
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
