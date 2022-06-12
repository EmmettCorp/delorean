package dialog

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

type Button struct {
	Text     string
	Callback func()
	active   bool
	coords   shared.Coords
}

func (b *Button) SetCoords(coords shared.Coords) {
	b.coords = coords
}

func (b *Button) GetCoords() shared.Coords {
	return b.coords
}

func (b *Button) OnClick(event tea.MouseMsg) error {
	b.Callback()

	return nil
}

func (b *Button) Render() string {
	if b.active {
		return activeButtonStyle.Render(b.Text)
	}

	return buttonStyle.Render(b.Text)
}
