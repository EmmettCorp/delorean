package gui

func (gui *Gui) informationStr() string {
	return gui.Config.GetVersion()
}

func (gui *Gui) handleInfoClick() error {
	if !gui.g.Mouse {
		return nil
	}
	return nil
}
