package shared

type clickableMock struct {
	coords Coords
}

func (cm *clickableMock) GetCoords() Coords {
	return cm.coords
}
func (cm *clickableMock) SetCoords(coords Coords) {
	cm.coords = coords
}

func (cm *clickableMock) OnClick(callback func() error) error {
	return nil
}
