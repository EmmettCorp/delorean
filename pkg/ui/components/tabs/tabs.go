package tabs

import (
	"errors"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
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
	currentTabID int
	tabs         []tab
}

func NewModel(titles []string) (Model, error) {
	if len(titles) == 0 {
		return Model{}, errors.New("empty titles")
	}

	m := Model{}

	x1 := 0
	for i := range titles {
		x2 := x1 + len(titles[i]) + 3
		m.tabs = append(m.tabs, tab{
			id:    i,
			title: titles[i],
			x1:    x1,
			y1:    0,
			x2:    x2,
			y2:    2,
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
		if m.currentTabID == i {
			tabs = append(tabs, activeTab.Render(m.tabs[i].title))
		} else {
			tabs = append(tabs, inactiveTab.Render((m.tabs[i].title)))
		}
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tabs...,
	)

	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	gap := tabGap.Render(strings.Repeat(" ", max(0, physicalWidth-lipgloss.Width(row)-2)))

	return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap))
}

func (m *Model) SetcurrentTabID(id int) {
	m.currentTabID = id
}

func (m *Model) OnClick(e tea.MouseMsg) int {
	for _, t := range m.tabs {
		if t.x1 <= e.X && e.X <= t.x2 && t.y1 <= e.Y && e.Y <= t.y2 {
			return t.id
		}
	}

	return m.currentTabID
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
