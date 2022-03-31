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

	for i := range titles {
		m.tabs = append(m.tabs, tab{
			id:    i,
			title: titles[i],
		})
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

	return lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
}

func (m *Model) SetcurrentTabID(id int) {
	m.currentTabID = id
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
