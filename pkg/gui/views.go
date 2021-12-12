package gui

import "github.com/jroimartin/gocui"

const (
	indent      = 0
	headerHight = 2
)

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
	gui.views = views{}
	gui.buttons.width = 0

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

	gui.views.status, err = gui.statusView(maxX, maxY)
	if err != nil {
		return err
	}

	if gui.views.snapshots, err = gui.g.SetView("snapshots", indent, headerHight, int(0.8*float32(maxX)), maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	gui.views.snapshots.Title = gui.views.snapshots.Name()
	if gui.views.schedule, err = gui.g.SetView("schedule", int(0.8*float32(maxX)), headerHight, maxX, maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	gui.views.schedule.Title = gui.views.schedule.Name()
	if gui.views.storage, err = gui.g.SetView("storage", indent, maxY-5, maxX, maxY); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	gui.views.storage.Title = gui.views.storage.Name()

	return nil
}

func (gui *Gui) allViews() []*gocui.View {
	return []*gocui.View{
		gui.views.snapshots,
		gui.views.schedule,
		gui.views.storage,
		gui.views.status,
	}
}
