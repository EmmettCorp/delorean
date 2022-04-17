package styles

import "github.com/charmbracelet/lipgloss"

// MainTextStyle is the general application text style.
// nolint:gochecknoglobals // global on purpose
var MainTextStyle = lipgloss.NewStyle().
	Foreground(DefaultTheme.MainText).
	Bold(true)
