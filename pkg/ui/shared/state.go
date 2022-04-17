/*
Package shared keeps shared domains.
*/
package shared

import (
	"errors"

	"github.com/EmmettCorp/delorean/pkg/config"
)

// State is the state of the ui application.
type State struct {
	ScreenHeight      int
	ScreenWidth       int
	CurrentTab        TabItem
	ClickableElements map[TabItem][]Clickable
	Config            *config.Config
	Areas             *uiAreas

	TestBool bool
}

func NewState(cfg *config.Config) *State {
	st := State{
		ClickableElements: make(map[TabItem][]Clickable),
		Config:            cfg,
		Areas:             initAreas(),
	}

	return &st
}

// Update changes the current tab.
func (s *State) Update(ti TabItem) {
	s.CurrentTab = ti
}

func (s *State) CleanClickable(ti TabItem) {
	s.ClickableElements[ti] = []Clickable{}
}

func (s *State) AppendClickable(ti TabItem, cc ...Clickable) error {
	for i := range cc {
		if err := validateClickable(cc[i]); err != nil {
			return err
		}
	}
	s.ClickableElements[ti] = append(s.ClickableElements[ti], cc...)

	return nil
}

func (s *State) FindClickable(x, y int) Clickable {
	var nearestClickable Clickable
	nearestCoords := Coords{}

	elements := make([]Clickable, 0, len(s.ClickableElements[AnyTab])+len(s.ClickableElements[s.CurrentTab]))
	elements = append(elements, s.ClickableElements[AnyTab]...)
	elements = append(elements, s.ClickableElements[s.CurrentTab]...)

	for i := range elements {
		coords := elements[i].GetCoords()
		if x >= coords.X1 && x <= coords.X2 && y >= coords.Y1 && y <= coords.Y2 {
			if nearestClickable == nil {
				nearestClickable = elements[i]
				nearestCoords = coords
			} else if x-coords.X1 < x-nearestCoords.X1 { // if there is a clickable inside another clickable
				nearestClickable = elements[i]
				nearestCoords = coords
			}
		}
	}

	return nearestClickable
}

func (s *State) ResizeAreas() {
	s.Areas.MainContent.Height = s.ScreenHeight - (s.Areas.TabBar.Height + s.Areas.HelpBar.Height)
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
