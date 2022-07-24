package shared

import (
	"errors"
)

// Clickable is a ui component that can be clicked.
type Clickable interface {
	GetCoords() Coords
	SetCoords(c Coords)
	SetCallback(callback func() error)
	OnClick() error
}

type ClickableItem struct {
	callback func() error
	coords   Coords
}

func (ci *ClickableItem) GetCoords() Coords {
	return ci.coords
}

func (ci *ClickableItem) SetCoords(c Coords) {
	ci.coords = c
}

func (ci *ClickableItem) SetCallback(callback func() error) {
	ci.callback = callback
}

func (ci *ClickableItem) OnClick() error {
	return ci.callback()
}

// Coords are coordinates of a clickable ui component.
type Coords struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
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

// ClickableComponent represents a logical component that help to identify on which component area clicked.
type ClickableComponent int

// Available components.
const (
	TabHeader ClickableComponent = iota
	SnapshotsButtonsBar
	SnapshotsList
	VolumesBar
	ScheduleBar
	HelpFooter
)

func getAllClickableComponents() []ClickableComponent {
	return []ClickableComponent{
		TabHeader,
		SnapshotsButtonsBar,
		SnapshotsList,
		VolumesBar,
		ScheduleBar,
		HelpFooter,
	}
}
