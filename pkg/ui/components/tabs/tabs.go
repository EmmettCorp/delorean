/*
Package tabs keeps the logic for tabs component.
*/
package tabs

import (
	"errors"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/ui/elements/tab"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	tabsLeftRightIndents = 2
)

type clickableTab interface {
	shared.Clickable
	GetTitle() string
	GetID() shared.TabItem
}

type Model struct {
	state *shared.State
	Tabs  []clickableTab
	keys  shared.KeyMap
}

func NewModel(state *shared.State, tabItems []shared.TabItem) (*Model, error) {
	if len(tabItems) == 0 {
		return nil, errors.New("empty tabItems")
	}

	m := Model{
		state: state,
		keys:  shared.GetKeyMaps(),
	}

	var x1 int
	for i := range tabItems {
		title := tabItems[i].String()
		x2 := x1 + len(title) + 3 // nolint:gomnd // 3 = 2 vertical bars + 1 space
		nt, err := tab.New(state, tabItems[i], shared.Coords{
			X1: x1,
			X2: x2,
			Y2: state.Areas.TabBar.Height,
		})
		if err != nil {
			return nil, err
		}
		m.Tabs = append(m.Tabs, nt)
		x1 = x2 + 1
	}

	return &m, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(msg, m.keys.Tab) {
			m.next()
		} else if key.Matches(msg, m.keys.ShiftTab) {
			m.prev()
		}
	}

	return m, nil
}

func (m *Model) View() string {
	var tabs []string
	for i := range m.Tabs {
		if m.state.CurrentTab == m.Tabs[i].GetID() {
			tabs = append(tabs, activeTab.Render(m.Tabs[i].GetTitle()))
		} else {
			tabs = append(tabs, inactiveTab.Render((m.Tabs[i].GetTitle())))
		}
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tabs...,
	)

	gap := tabGap.Render(strings.Repeat(" ", max(0, m.state.ScreenWidth-lipgloss.Width(row)-tabsLeftRightIndents)))

	return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap))
}

func (m *Model) next() {
	i := m.getNextTabIndex()
	m.state.CurrentTab = m.Tabs[i].GetID()
}

func (m *Model) prev() {
	i := m.getPrevTabIndex()
	m.state.CurrentTab = m.Tabs[i].GetID()
}

func (m *Model) getCurrentTabIndex() int {
	for i := range m.Tabs {
		if m.state.CurrentTab == m.Tabs[i].GetID() {
			return i
		}
	}

	return 0
}

func (m *Model) getNextTabIndex() int {
	return (m.getCurrentTabIndex() + 1) % len(m.Tabs)
}

func (m *Model) getPrevTabIndex() int {
	return (m.getCurrentTabIndex() - 1 + len(m.Tabs)) % len(m.Tabs)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
