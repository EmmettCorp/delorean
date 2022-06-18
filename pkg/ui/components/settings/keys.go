package settings

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Enter key.Binding
}

func getKeyMaps() keyMap {
	return keyMap{
		Enter: key.NewBinding(
			key.WithKeys(tea.KeyEnter.String()),
			key.WithHelp("Enter", "activate/deactivate subvolume"),
		),
	}
}
