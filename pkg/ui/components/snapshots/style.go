package snapshots

import "github.com/charmbracelet/lipgloss"

// nolint:gochecknoglobals // this is used only in this package
var (
	docStyle = lipgloss.NewStyle().Margin(0, 0)
	subtle   = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	inactive = lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#5C5C5C"}
)
