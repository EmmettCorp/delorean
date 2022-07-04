/*
Package alert keeps helpers to create a standard alert window.
*/
package alert

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	dWidth = 30
)

// Model is a dialog model.
type Model struct {
	title string
}

// New creates and returns a new dialog model.
func New() *Model {
	return &Model{}
}

func (m *Model) View(w, h int) string {
	ui := lipgloss.NewStyle().Width(dWidth).Align(lipgloss.Center).Render(m.title)

	return lipgloss.Place(w, h,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(subtle),
	)
}

func (m *Model) SetTitle(t string) {
	m.title = t
}
