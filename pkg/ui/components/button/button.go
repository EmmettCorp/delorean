/*
Package button keeps all the logic for button component.
*/
package button

import "github.com/EmmettCorp/delorean/pkg/ui/shared"

// Model here represents the button model.
type Model interface {
	shared.Clickable
	SetTitle(title string)
	GetTitle() string
}

func DrawButton(title string) string {
	return buttonWithBorder.Render(title)
}
