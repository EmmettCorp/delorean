/*
Package help keeps the logic for bottom help bar.
*/
package help

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/styles"
	"github.com/charmbracelet/lipgloss"
)

// nolint:gochecknoglobals // this is used only in this package
var (
	helpTextStyle = lipgloss.NewStyle().Foreground(styles.DefaultTheme.InactiveText)
	helpStyle     = lipgloss.NewStyle().
			Height(shared.HelpBarHeight).
			BorderForeground(styles.DefaultTheme.InactiveText)
)
