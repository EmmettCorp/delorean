package help

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	FooterHeight = 2

	helpTextStyle = lipgloss.NewStyle().Foreground(styles.DefaultTheme.InactiveText)
	helpStyle     = lipgloss.NewStyle().
			Height(FooterHeight - 1).
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(styles.DefaultTheme.InactiveText)
)
