package gui

import "github.com/jroimartin/gocui"

type views struct {
	snapshots  *gocui.View
	schedule   *gocui.View
	storage    *gocui.View
	status     *gocui.View
	createBtn  *gocui.View
	restoreBtn *gocui.View
	deleteBtn  *gocui.View
}

func (gui *Gui) layout(g *gocui.Gui) error {
	var err error
	maxX, maxY := gui.g.Size()

	// draw buttons: order matters
	gui.views.createBtn, err = gui.createButton()
	if err != nil {
		return err
	}
	gui.views.restoreBtn, err = gui.restoreButton()
	if err != nil {
		return err
	}
	gui.views.deleteBtn, err = gui.deleteButton()
	if err != nil {
		return err
	}

	// draw pannels
	gui.views.status, err = gui.statusView(maxX, maxY)
	if err != nil {
		return err
	}
	gui.views.snapshots, err = gui.snapshotsView(maxX, maxY)
	if err != nil {
		return err
	}
	gui.views.schedule, err = gui.scheduleView(maxX, maxY)
	if err != nil {
		return err
	}
	gui.views.storage, err = gui.storageView(maxX, maxY)
	if err != nil {
		return err
	}

	return nil
}

func (gui *Gui) allViews() []*gocui.View {
	return []*gocui.View{
		gui.views.snapshots,
		gui.views.schedule,
		gui.views.storage,
		gui.views.status,
		gui.views.createBtn,
		gui.views.restoreBtn,
		gui.views.deleteBtn,
	}
}
