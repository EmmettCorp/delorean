package shared

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

// Successor is a component that can be added to a Clickable.
type Successor interface {
	OnClick(event tea.MouseMsg)
	GetCoords() (int, int, int, int) // x1, y1, x2, y2
}

// Clickable is a ui component that can be clicked.
type Clickable struct {
	Successors   []Successor
	X1           int
	Y1           int
	X2           int
	Y2           int
	ClickHandler func(event tea.MouseMsg)
}

// AddSuccessor adds a successor to the Clickable.
func (cl *Clickable) AddSuccessor(s Successor) error {
	err := cl.validate()
	if err != nil {
		return err
	}

	cl.Successors = append(cl.Successors, s)

	return nil
}

// OnClick is called when the user clicks on a Clickable.
// This function should be called by the Clickable's parent.
// The main ancestor is the app that has coordinates x1 = 0, y1 = 0, x2 = max, y2 = max.
func (cl *Clickable) OnClick(event tea.MouseMsg) {
	// first check if the click is within any of the successors
	for _, s := range cl.Successors {
		x1, y1, x2, y2 := s.GetCoords()
		if x1 <= event.X && event.X <= x2 && y1 <= event.Y && event.Y <= y2 {
			s.OnClick(event)
			return
		}
	}

	// if no successor was clicked, call the ClickHandler function of the Clickable if set
	if cl.ClickHandler != nil {
		cl.ClickHandler(event)
	}
}

func (cl *Clickable) validate() error {
	for _, e := range cl.Successors {
		x1, y1, x2, y2 := e.GetCoords()
		if x1 >= x2 || y1 >= y2 {
			return errors.New("invalid element coordinates")
		}
	}

	return nil
}
