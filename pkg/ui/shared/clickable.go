package shared

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

// Coords are coordinates of a clickable ui component.
type Coords struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// Clickable is a ui component that can be clicked.
type Clickable interface {
	GetCoords() Coords
	SetCoords(c Coords)
	OnClick(event tea.MouseMsg) error
}

func validateClickable(c Clickable) error {
	if c == nil {
		return errors.New("nil clickable")
	}

	coords := c.GetCoords()
	if coords.X1 >= coords.X2 || coords.Y1 >= coords.Y2 {
		return errors.New("invalid coords")
	}

	return nil
}
