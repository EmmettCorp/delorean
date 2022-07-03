package shared

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Tab      key.Binding
	ShiftTab key.Binding
	Up       key.Binding
	Down     key.Binding
	PrevPage key.Binding
	NextPage key.Binding
	Create   key.Binding
	Restore  key.Binding
	Delete   key.Binding
	Volume   key.Binding
	Help     key.Binding
	Quit     key.Binding
}

func (k KeyMap) shortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Tab}
}

func (k KeyMap) navigation() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.PrevPage, k.NextPage,
	}
}

func (k KeyMap) SettingsHelp() [][]key.Binding {
	return [][]key.Binding{
		k.shortHelp(),
		k.navigation(),
		{k.Volume},
	}
}

func (k KeyMap) SnapshotsHelp() [][]key.Binding {
	return [][]key.Binding{
		k.shortHelp(),
		k.navigation(),
		{k.Create, k.Restore, k.Delete},
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Tab}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit, k.Tab},
		{k.Up, k.Down, k.PrevPage, k.NextPage},
	}
}

// GetKeyMaps returns all the shortcats available.
func GetKeyMaps() KeyMap {
	return KeyMap{
		Tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next tab"),
		),
		ShiftTab: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("Shift+tab", "previous tab"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		PrevPage: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "prev page"),
		),
		NextPage: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "next page"),
		),
		Create: key.NewBinding(
			key.WithKeys(tea.KeyCtrlCaret.String()),
			key.WithHelp("Ctrl+Enter", "create"),
		),
		Restore: key.NewBinding(
			key.WithKeys(tea.KeyEnter.String()),
			key.WithHelp("Enter/r", "restore"),
		),
		Delete: key.NewBinding(
			key.WithKeys(tea.KeyDelete.String(), "d"),
			key.WithHelp("Delete/d", "delete snapshot"),
		),
		Volume: key.NewBinding(
			key.WithKeys(tea.KeyEnter.String()),
			key.WithHelp("Enter", "toggle activate subvolume"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}
