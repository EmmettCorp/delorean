package gui

import (
	"time"

	"github.com/EmmettCorp/delorean/pkg/rate"
	"github.com/jroimartin/gocui"
)

const (
	minWidthConst  = 57
	minHeightConst = 15
)

type view struct {
	name    string
	x0      int
	x1      int
	y0      int
	y1      int
	limiter *rate.Limiter
}

type views struct {
	createBtn  view
	restoreBtn view
	deleteBtn  view
	status     view
	snapshots  view
	schedule   view
	volumes    view
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
	// |-volumes---------------------------------------------------------------------|
	// |                                                                             |
	// | here is volumes data                                                        |
	// |                                                                             |
	// -------------------------------------------------------------------------------

	// so the order below matters

	gui.views.createBtn.name = "create"
	gui.views.createBtn.x0 = 0
	gui.views.createBtn.x1 = gui.views.createBtn.x0 + len(gui.views.createBtn.name) + 1
	gui.views.createBtn.y0 = headerY0
	gui.views.createBtn.y1 = headerY1
	// Limiter for create button is needed to allow to finish create snapshot operation.
	// There is no real point in real life doing snapshots every second.
	// If allow user to call btrfs.CreateSnapshot several times a second it could cause a exec.Command call error.
	gui.views.createBtn.limiter = rate.NewLimiter(time.Second * 2)

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

	gui.views.volumes.name = "volumes"
	gui.views.volumes.x0 = 0
	gui.views.volumes.x1 = gui.maxX
	gui.views.volumes.y0 = gui.views.snapshots.y1
	gui.views.volumes.y1 = gui.maxY

	gui.views.errorView.name = "error"
	gui.views.errorView.x0 = 0
	gui.views.errorView.x1 = gui.maxX - 1
	gui.views.errorView.y0 = 0
	gui.views.errorView.y1 = gui.maxY - 1
}

func (gui *Gui) layout(g *gocui.Gui) error {
	var err error
	gui.maxX, gui.maxY = gui.g.Size()

	if gui.maxX < minWidthConst || gui.maxY < minHeightConst {
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
	_, err = gui.volumesView()
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
		gui.views.volumes.name,
	} {
		v, err := gui.g.View(name)
		if err != nil {
			continue
		}
		vv = append(vv, v)
	}

	return vv
}
