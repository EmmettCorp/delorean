package button

import "github.com/charmbracelet/lipgloss"

// nolint:gochecknoglobals // this is used only in tabs package
var (
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#AD58B4"}

	buttonBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	buttonWithBorder = lipgloss.NewStyle().
				Border(buttonBorder, true).
				BorderForeground(highlight).
				Padding(0, 1)
)
