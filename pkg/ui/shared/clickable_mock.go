package shared

import tea "github.com/charmbracelet/bubbletea"

type clickableMock struct {
	coords Coords
}

func (cm *clickableMock) GetCoords() Coords {
	return cm.coords
}
func (cm *clickableMock) SetCoords(coords Coords) {
	cm.coords = coords
}

func (cm *clickableMock) OnClick(event tea.MouseMsg) error {
	return nil
}
