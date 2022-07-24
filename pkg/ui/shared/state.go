/*
Package shared keeps shared domains.
*/
package shared

import (
	"github.com/EmmettCorp/delorean/pkg/config"
)

// State is the state of the ui application.
type State struct {
	ScreenHeight      int
	ScreenWidth       int
	CurrentTab        TabItem
	ClickableElements map[ClickableComponent][]Clickable
	Config            *config.Config
	Areas             *uiAreas
	UpdateSnapshots   bool
}

func NewState(cfg *config.Config) *State {
	st := State{
		ClickableElements: make(map[ClickableComponent][]Clickable),
		Config:            cfg,
		Areas:             initAreas(),
	}

	return &st
}

// Update changes the current tab.
func (s *State) Update(ti TabItem) {
	s.CurrentTab = ti
}

func (s *State) CleanClickable(comp ClickableComponent) {
	s.ClickableElements[comp] = []Clickable{}
}

func (s *State) AppendClickable(comp ClickableComponent, cc ...Clickable) error {
	for i := range cc {
		if err := validateClickable(cc[i]); err != nil {
			return err
		}
	}
	s.ClickableElements[comp] = append(s.ClickableElements[comp], cc...)

	return nil
}

func (s *State) FindClickable(x, y int) Clickable {
	var nearestClickable Clickable

	elements := s.getAvailableClickable()

	nearestCoords := Coords{}
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
	s.Areas.MainContent.Width = s.ScreenWidth
	s.Areas.TabBar.Width = s.ScreenWidth
	s.Areas.HelpBar.Width = s.ScreenWidth
}

func (s *State) getAvailableClickable() []Clickable {
	elements := []Clickable{}
	elements = append(elements, s.ClickableElements[TabHeader]...)
	if s.CurrentTab == SnapshotsTab {
		elements = append(elements, s.ClickableElements[SnapshotsButtonsBar]...)
		elements = append(elements, s.ClickableElements[SnapshotsList]...)
	} else if s.CurrentTab == SettingsTab {
		elements = append(elements, s.ClickableElements[VolumesBar]...)
		elements = append(elements, s.ClickableElements[ScheduleBar]...)
	}

	return elements
}

func (s *State) GetActiveVolumesIDs() []string {
	ids := []string{}
	for i := range s.Config.Volumes {
		if s.Config.Volumes[i].Active {
			ids = append(ids, s.Config.Volumes[i].ID)
		}
	}

	return ids
}
