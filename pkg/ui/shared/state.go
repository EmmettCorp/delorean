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
	ClickableElements map[TabItem][]Clickable
}

func (s *State) Update(ti TabItem) {
	s.CurrentTab = ti
}

func (s *State) CleanClickable(ti TabItem) {
	s.ClickableElements[ti] = []Clickable{}
}

func (s *State) AppendClickable(ti TabItem, c ...Clickable) {
	s.ClickableElements[ti] = append(s.ClickableElements[ti], c...)
}

func (s *State) FindClickable(x, y int) Clickable {
	var nearestClickable Clickable
	nearestCoords := Coords{}

	elements := append(s.ClickableElements[AnyTab], s.ClickableElements[s.CurrentTab]...)

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
