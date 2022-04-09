package shared

import (
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
	OnClick(event tea.MouseMsg)
}
