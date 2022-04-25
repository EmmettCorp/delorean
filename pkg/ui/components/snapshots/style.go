package snapshots

import (
	"github.com/charmbracelet/lipgloss"
)

// nolint:gochecknoglobals // this is used only in this package
var (
	docStyle  = lipgloss.NewStyle().Margin(0, 0)
	itemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
			Padding(0, 0, 0, 2)
	selectedItemStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
				Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
				Padding(0, 0, 0, 1)
)
