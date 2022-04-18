/*
Package styles keeps default styles.
*/
package styles

import "github.com/charmbracelet/lipgloss"

// nolint:gochecknoglobals // this is used only in this package
var indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#383B5B"}

// Theme is a struct to keep all the application styles.
type Theme struct {
	MainText           lipgloss.AdaptiveColor
	BrightMainText     lipgloss.AdaptiveColor
	Border             lipgloss.AdaptiveColor
	SecondaryBorder    lipgloss.AdaptiveColor
	WarningText        lipgloss.AdaptiveColor
	SuccessText        lipgloss.AdaptiveColor
	FaintBorder        lipgloss.AdaptiveColor
	FaintText          lipgloss.AdaptiveColor
	SelectedBackground lipgloss.AdaptiveColor
	SecondaryText      lipgloss.AdaptiveColor
	SubleMainText      lipgloss.AdaptiveColor
	InactiveText       lipgloss.AdaptiveColor
}

// nolint:gochecknoglobals // this is used only in this package
var subtleIndigo = lipgloss.AdaptiveColor{Light: "#5A57B5", Dark: "#242347"}

// DefaultTheme is an application default theme.
// nolint:gochecknoglobals // global on purpose
var DefaultTheme = Theme{
	MainText:           lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"},
	BrightMainText:     lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"},
	SubleMainText:      subtleIndigo,
	Border:             lipgloss.AdaptiveColor{Light: indigo.Light, Dark: indigo.Dark},
	SecondaryBorder:    lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#39386b"},
	WarningText:        lipgloss.AdaptiveColor{Light: "#F23D5C", Dark: "#F23D5C"},
	SuccessText:        lipgloss.AdaptiveColor{Light: "#3DF294", Dark: "#3DF294"},
	FaintBorder:        lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#2b2b40"},
	FaintText:          lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#3E4057"},
	SelectedBackground: lipgloss.AdaptiveColor{Light: subtleIndigo.Light, Dark: "#39386b"},
	SecondaryText:      lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#666CA6"},
	InactiveText:       lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#5C5C5C"},
}

const (
	SingleRuneWidth    = 4
	MainContentPadding = 1
)
