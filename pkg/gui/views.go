package gui

import (
	"github.com/jroimartin/gocui"
)

const (
	MIN_WIDTH  = 57
	MIN_HEIGHT = 15
)

type view struct {
	name string
	x0   int
	x1   int
	y0   int
	y1   int
}

type views struct {
	createBtn  view
	restoreBtn view
	deleteBtn  view
	status     view
	snapshots  view
	schedule   view
	storage    view
	errorView  view
}

func (gui *Gui) initViews() {
	gui.maxX, gui.maxY = gui.g.Size()
	indent := 1
	headerY0 := -1
	headerY1 := 1

	// this is how ui is supposed to look like
	// -------------------------------------------------------------------------------
	// | create | restore | delete | status bar                                      |
	// |-----------------------------------------------------------------------------|
	// |-snapshots----------------------------------------|-schedule-----------------|
	// |                                                  |                          |
	// | here is a snapshots list                         | schedule list            |
	// |                                                  |                          |
	// |                                                  |                          |
	// ...                                               ...                        ...
	// |                                                  |                          |
	// |-storage---------------------------------------------------------------------|
	// |                                                                             |
	// | here is storage data                                                        |
	// |                                                                             |
	// -------------------------------------------------------------------------------

	// so the order below matters

	gui.views.createBtn.name = "create"
	gui.views.createBtn.x0 = 0
	gui.views.createBtn.x1 = gui.views.createBtn.x0 + len(gui.views.createBtn.name) + 1
	gui.views.createBtn.y0 = headerY0
	gui.views.createBtn.y1 = headerY1

	gui.views.restoreBtn.name = "restore"
	gui.views.restoreBtn.x0 = gui.views.createBtn.x1 + indent
	gui.views.restoreBtn.x1 = gui.views.restoreBtn.x0 + len(gui.views.restoreBtn.name) + 1
	gui.views.restoreBtn.y0 = headerY0
	gui.views.restoreBtn.y1 = headerY1

	gui.views.deleteBtn.name = "delete"
	gui.views.deleteBtn.x0 = gui.views.restoreBtn.x1 + indent
	gui.views.deleteBtn.x1 = gui.views.deleteBtn.x0 + len(gui.views.deleteBtn.name) + 1
	gui.views.deleteBtn.y0 = headerY0
	gui.views.deleteBtn.y1 = headerY1

	gui.views.status.name = "status"
	gui.views.status.x0 = gui.views.deleteBtn.x1 + indent
	gui.views.status.x1 = gui.maxX
	gui.views.status.y0 = headerY0
	gui.views.status.y1 = headerY1

	gui.views.snapshots.name = "snapshots"
	gui.views.snapshots.x0 = 0
	gui.views.snapshots.x1 = int(0.8 * float32(gui.maxX))
	gui.views.snapshots.y0 = headerY1 + 1
	gui.views.snapshots.y1 = gui.maxY - 5

	gui.views.schedule.name = "schedule"
	gui.views.schedule.x0 = gui.views.snapshots.x1
	gui.views.schedule.x1 = gui.maxX
	gui.views.schedule.y0 = headerY1 + 1
	gui.views.schedule.y1 = gui.maxY - 5

	gui.views.storage.name = "storage"
	gui.views.storage.x0 = 0
	gui.views.storage.x1 = gui.maxX
	gui.views.storage.y0 = gui.views.snapshots.y1
	gui.views.storage.y1 = gui.maxY

	gui.views.errorView.name = "error"
	gui.views.storage.x0 = 0
	gui.views.storage.x1 = gui.maxX - 1
	gui.views.storage.y0 = 0
	gui.views.storage.y1 = gui.maxY - 1
}

func (gui *Gui) layout(g *gocui.Gui) error {
	var err error
	gui.maxX, gui.maxY = gui.g.Size()

	if gui.maxX < MIN_WIDTH || gui.maxY < MIN_HEIGHT {
		_, err = gui.errorView()
		return err
	} else {
		err := gui.deleteErrorView()
		if err != nil {
			return err
		}
	}

	_, err = gui.createButton()
	if err != nil {
		return err
	}
	_, err = gui.restoreButton()
	if err != nil {
		return err
	}
	_, err = gui.deleteButton()
	if err != nil {
		return err
	}

	_, err = gui.statusView()
	if err != nil {
		return err
	}
	_, err = gui.snapshotsView()
	if err != nil {
		return err
	}
	_, err = gui.scheduleView()
	if err != nil {
		return err
	}
	_, err = gui.storageView()
	if err != nil {
		return err
	}

	return nil
}

func (gui *Gui) allViews() []*gocui.View {
	vv := []*gocui.View{}
	for _, name := range []string{
		gui.views.createBtn.name,
		gui.views.restoreBtn.name,
		gui.views.deleteBtn.name,
		gui.views.status.name,
		gui.views.snapshots.name,
		gui.views.schedule.name,
		gui.views.storage.name,
	} {
		v, err := gui.g.View(name)
		if err != nil {
			continue
		}
		vv = append(vv, v)
	}

	return vv
}
