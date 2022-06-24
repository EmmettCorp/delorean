package settings

import (
	"github.com/charmbracelet/lipgloss"
)

// nolint:gochecknoglobals // used only in this package
var (
	subtle            = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	signActiveStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#6aa84f"))
	signInactiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#c32a2a"))
	listStyle         = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, true, false, false).
				BorderForeground(subtle).
				MarginRight(2).PaddingRight(2)

	normalItem = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).PaddingLeft(2)

	selectedItem = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).PaddingLeft(1)
)
