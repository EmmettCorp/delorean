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
