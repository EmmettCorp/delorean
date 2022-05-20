package snapshots

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Delete key.Binding
}

func getKeyMaps() keyMap {
	return keyMap{
		Delete: key.NewBinding(
			key.WithKeys(tea.KeyDelete.String(), "d"),
			key.WithHelp("Delete", "delete snapshot"),
		),
	}
}
