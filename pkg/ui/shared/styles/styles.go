package styles

import "github.com/charmbracelet/lipgloss"

// MainTextStyle is the general application text style.
// nolint:gochecknoglobals // global on purpose
var MainDocStyle = lipgloss.NewStyle()

var MainTextStyle = MainDocStyle.
	Foreground(DefaultTheme.MainText).
	Bold(true)
