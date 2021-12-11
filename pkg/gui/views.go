package gui

import "github.com/jroimartin/gocui"

const (
	indent      = 0
	headerHight = 2
)

type views struct {
	snapshots *gocui.View
	path      *gocui.View
	termital  *gocui.View
	status    *gocui.View
}

func (gui *Gui) initViews() error {
	var err error
	maxX, maxY := gui.g.Size()
	gui.views = views{}

	if gui.views.status, err = gui.g.SetView("status", gui.buttons.width, -1, maxX, headerHight-1); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	if gui.views.snapshots, err = gui.g.SetView("snapshots", indent, headerHight, int(0.8*float32(maxX)), maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	gui.views.snapshots.Title = gui.views.snapshots.Name()
	if gui.views.path, err = gui.g.SetView("path", int(0.8*float32(maxX)), headerHight, maxX, maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	gui.views.path.Title = gui.views.path.Name()
	if gui.views.termital, err = gui.g.SetView("terminal", indent, maxY-5, maxX, maxY); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	gui.views.termital.Title = gui.views.termital.Name()

	return nil
}
