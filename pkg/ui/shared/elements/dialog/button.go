package dialog

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
)

type Button struct {
	shared.ClickableItem
	Text   string
	active bool
}

func (b *Button) Render() string {
	if b.active {
		return activeButtonStyle.Render(b.Text)
	}

	return buttonStyle.Render(b.Text)
}
