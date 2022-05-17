package styles

import "github.com/charmbracelet/lipgloss"

// MainStyle is the general application text style.
// nolint:gochecknoglobals // global on purpose
var (
	MainDocStyle = lipgloss.NewStyle()

	MainTextStyle = MainDocStyle.
			Foreground(DefaultTheme.MainText).
			Bold(true)
)
