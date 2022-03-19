package prssection

import (
	"github.com/EmmettCorp/delorean/pkg/ui/components/table"
	"github.com/EmmettCorp/delorean/pkg/utils"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Id        int
	spinner   spinner.Model
	isLoading bool
	table     table.Model
}

func (m *Model) View() string {
	var spinnerText *string
	if m.isLoading {
		spinnerText = utils.StringPtr(lipgloss.JoinHorizontal(lipgloss.Top,
			spinnerStyle.Copy().Render(m.spinner.View()),
			"Fetching Pull Requests...",
		))
	}

	return containerStyle.Copy().Render(
		m.table.View(spinnerText),
	)
}
