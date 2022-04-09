package divider

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

func Dot(ac lipgloss.AdaptiveColor) string {
	if ac.Light == "" || ac.Dark == "" {
		ac = subtle
	}

	return lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(ac).
		String()
}

func Horizontal(width int, ac lipgloss.AdaptiveColor) string {
	if ac.Light == "" || ac.Dark == "" {
		ac = subtle
	}

	div := lipgloss.NewStyle().
		SetString("─").
		Foreground(ac).
		String()

	return lipgloss.JoinHorizontal(lipgloss.Bottom, strings.Repeat(div, width))
}
