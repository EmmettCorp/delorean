package listviewport

import (
	"github.com/EmmettCorp/delorean/pkg/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	pagerHeight = 2

	pagerStyle = lipgloss.NewStyle().
			Height(pagerHeight).
			MaxHeight(pagerHeight).
			PaddingTop(1).
			Bold(true).
			Foreground(styles.DefaultTheme.FaintText)
)
