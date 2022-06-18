/*
Package styles keeps default styles.
*/
package styles

import "github.com/charmbracelet/lipgloss"

// Theme is a struct to keep all the application styles.
type Theme struct {
	MainText     lipgloss.AdaptiveColor
	InactiveText lipgloss.AdaptiveColor
}

// DefaultTheme is an application default theme.
// nolint:gochecknoglobals // global on purpose
var DefaultTheme = Theme{
	MainText:     lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"},
	InactiveText: lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#5C5C5C"},
}
