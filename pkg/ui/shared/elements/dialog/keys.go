package dialog

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Left  key.Binding
	Right key.Binding
	Enter key.Binding
}

func getKeyMaps() keyMap {
	return keyMap{
		Left: key.NewBinding(
			key.WithKeys(tea.KeyLeft.String(), "h"),
			key.WithHelp("Left", "left"),
		),
		Right: key.NewBinding(
			key.WithKeys(tea.KeyRight.String(), "l"),
			key.WithHelp("Right", "right"),
		),
		Enter: key.NewBinding(
			key.WithKeys(tea.KeyEnter.String()),
			key.WithHelp("Enter", "enter"),
		),
	}
}
