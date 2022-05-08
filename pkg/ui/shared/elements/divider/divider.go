/*
Package divider keeps helpers to draw dividers.
*/
package divider

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Dot is a dot divider.
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

// HorizontalLine is a horizontal line.
func HorizontalLine(width int, ac lipgloss.AdaptiveColor) string {
	if ac.Light == "" || ac.Dark == "" {
		ac = subtle
	}

	div := lipgloss.NewStyle().
		SetString("─").
		Foreground(ac).
		String()

	return lipgloss.JoinHorizontal(lipgloss.Bottom, strings.Repeat(div, width))
}
