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

type tab struct {
	id    int
	title string
	x1    int
	y1    int
	x2    int
	y2    int
}

type Model struct {
	state *shared.State
	tabs  []tab
}

func NewModel(state *shared.State, titles []string) (Model, error) {
	if len(titles) == 0 {
		return Model{}, errors.New("empty titles")
	}

	m := Model{
		state: state,
	}

	var x1 int
	for i := range titles {
		x2 := x1 + len(titles[i]) + 3 // nolint:gomnd // 3 = 2 vertical bars + 1 space
		m.tabs = append(m.tabs, tab{
			id:    i,
			title: titles[i],
			x1:    x1,
			x2:    x2,
			y2:    TabsHeigh,
		})
		x1 = x2 + 1
	}

	return m, nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var tabs []string
	for i := range m.tabs {
		if m.state.CurrentTab == i {
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

func (m *Model) OnClick(e tea.MouseMsg) {
	for _, t := range m.tabs {
		if t.x1 <= e.X && e.X <= t.x2 && t.y1 <= e.Y && e.Y <= t.y2 {
			m.state.CurrentTab = t.id

			return
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
