package tabs

import (
	"github.com/EmmettCorp/delorean/pkg/ui/context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	CurrSectionId int
}

func NewModel() Model {
	return Model{
		CurrSectionId: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View(ctx context.ProgramContext) string {
	sectionTitles := []string{
		"Snapshots",
		"Schedule",
		"Settings",
	}

	var tabs []string
	for i, sectionTitle := range sectionTitles {
		if m.CurrSectionId == i {
			tabs = append(tabs, activeTab.Render(sectionTitle))
		} else {
			tabs = append(tabs, tab.Render(sectionTitle))
		}
	}

	return tabsRow.Copy().
		Width(ctx.ScreenWidth).
		MaxWidth(ctx.ScreenWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))
}

func (m *Model) SetCurrSectionId(id int) {
	m.CurrSectionId = id
}
