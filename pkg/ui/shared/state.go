/*
Package shared keeps shared domains.
*/
package shared

import (
	"github.com/EmmettCorp/delorean/pkg/domain"
)

// State is the state of the ui application.
type State struct {
	ActiveVolumes     []domain.Volume
	CurrentTab        TabItem
	ClickableElements []Clickable
}

func (s *State) UpdateClickable(cc []Clickable) {
	s.ClickableElements = cc
}

func (s *State) CleanClickable() {
	s.ClickableElements = []Clickable{}
}

func (s *State) AppendClickable(c ...Clickable) {
	s.ClickableElements = append(s.ClickableElements, c...)
}

func (s *State) FindClickable(x, y int) Clickable {
	var nearestClickable Clickable
	nearestCoords := Coords{}

	for i := range s.ClickableElements {
		coords := s.ClickableElements[i].GetCoords()
		if x >= coords.X1 && x <= coords.X2 && y >= coords.Y1 && y <= coords.Y2 {
			if nearestClickable == nil {
				nearestClickable = s.ClickableElements[i]
				nearestCoords = coords
			} else { // if there is a clickable inside another clickable
				if coords.X2-coords.X1 < nearestCoords.X2-nearestCoords.X1 {
					nearestClickable = s.ClickableElements[i]
					nearestCoords = coords
				}
			}
		}
	}

	return nearestClickable
}
