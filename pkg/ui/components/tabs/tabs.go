/*
Package tabs keeps the logic for tabs component.
*/
package tabs

import (
	"errors"

	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/tab"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type clickableTab interface {
	shared.Clickable
	GetID() shared.TabItem
	Render() string
}

type Model struct {
	state  *shared.State
	Tabs   []clickableTab
	keys   shared.KeyMap
	coords shared.Coords
}

// New creates a new tab model and set the pointer back to the callee.
func New(state *shared.State, tabItems []shared.TabItem) (*Model, error) {
	if len(tabItems) == 0 {
		return nil, errors.New("empty tabItems")
	}

	m := Model{
		state: state,
		keys:  shared.GetKeyMaps(),
		coords: shared.Coords{
			X1: 0,
			Y1: 0,
			X2: state.Areas.TabBar.Height,
			Y2: state.Areas.TabBar.Width,
		},
	}

	var x1 int
	for i := range tabItems {
		title := tabItems[i].String()
		x2 := x1 + lipgloss.Width(title) + 3 // nolint:gomnd // 3 = 2 vertical bars + 1 space
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
	tabs := make([]string, 0, len(m.Tabs))
	for i := range m.Tabs {
		tabs = append(tabs, m.Tabs[i].Render())
	}

	return tab.RenderTabBar(m.state.ScreenWidth, tabs)
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

func (m *Model) GetCoords() shared.Coords {
	return m.coords
}
