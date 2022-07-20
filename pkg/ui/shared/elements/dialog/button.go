package dialog

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
)

type Button struct {
	Text     string
	callback func() error
	active   bool
	coords   shared.Coords
}

func (b *Button) SetCoords(coords shared.Coords) {
	b.coords = coords
}

func (b *Button) GetCoords() shared.Coords {
	return b.coords
}

func (b *Button) SetCallback(callback func() error) {
	b.callback = callback
}

func (b *Button) OnClick() error {
	return b.callback()
}

func (b *Button) Render() string {
	if b.active {
		return activeButtonStyle.Render(b.Text)
	}

	return buttonStyle.Render(b.Text)
}
