package snapshots

import (
	"github.com/charmbracelet/lipgloss"
)

// nolint:gochecknoglobals // this is used only in this package
var (
	docStyle          = lipgloss.NewStyle().Margin(0, 0)
	itemStyle         = lipgloss.NewStyle().BorderLeft(true).Margin(0, 0, 1, 1)
	selectedItemStyle = lipgloss.NewStyle().BorderLeft(true).Margin(0, 0, 1, 1).Foreground(lipgloss.Color("170"))
)
