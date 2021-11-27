package gui

func (gui *Gui) isPopupPanel(viewName string) bool {
	return viewName == "credentials" || viewName == "confirmation" || viewName == "menu"
}
