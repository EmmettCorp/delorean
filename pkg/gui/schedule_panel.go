package gui

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
)

const maxSnapshotAmount = 99
const maxScheduleItems = 4

type scheduleEditor struct {
	g *Gui
}

func (gui *Gui) scheduleView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.schedule.name,
		int(0.8*float32(gui.maxX)),
		gui.views.schedule.y0,
		gui.maxX,
		gui.maxY-5,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.schedule.name, err)
			return nil, err
		}

		view.Title = gui.views.schedule.name
		view.Editor = &scheduleEditor{
			g: gui,
		}

		gui.drawSchedule(view)

		err := gui.g.SetKeybinding(gui.views.schedule.name, gocui.MouseLeft, gocui.ModNone, gui.editSchedule)
		if err != nil {
			return nil, err
		}
		err = gui.g.SetKeybinding(gui.views.schedule.name, gocui.KeyEsc, gocui.ModNone, gui.escapeSchedule)
		if err != nil {
			return nil, err
		}
		err = gui.g.SetKeybinding(gui.views.schedule.name, gocui.KeyEnter, gocui.ModNone, gui.saveSchedule)
		if err != nil {
			return nil, err
		}

	}

	return view, nil
}

func (gui *Gui) drawSchedule(view *gocui.View) {
	fmt.Fprintf(view, " Monthly:%2d\n", gui.config.Schedule.Monthly)
	fmt.Fprintf(view, " Weekly: %2d\n", gui.config.Schedule.Weekly)
	fmt.Fprintf(view, " Daily:  %2d\n", gui.config.Schedule.Daily)
	fmt.Fprintf(view, " Hourly: %2d\n", gui.config.Schedule.Hourly)
	fmt.Fprintf(view, " Boot:   %2d\n", gui.config.Schedule.Boot)
}

func (gui *Gui) saveSchedule(g *gocui.Gui, view *gocui.View) error {
	err := gui.config.Save()
	if err != nil {
		return fmt.Errorf("can't save config: %v", err)
	}

	gui.state.status = " schedule is saved "
	return nil
}

func (gui *Gui) editSchedule(g *gocui.Gui, view *gocui.View) error {
	gui.g.Cursor = false
	view.Highlight = true
	view.Editable = true
	view.SelBgColor = gocui.ColorGreen

	_, err := gui.g.SetCurrentView(gui.views.schedule.name)
	if err != nil {
		return err
	}

	// _, err = gui.g.SetViewOnTop(gui.views.schedule.name)

	return err
}

func (gui *Gui) escapeSchedule(g *gocui.Gui, view *gocui.View) error {
	view.Highlight = false
	view.SelBgColor = gocui.ColorDefault

	gui.setDefaultStatus()
	gui.g.SetCurrentView(gui.views.status.name)
	return nil
}

func (gui *Gui) updateSchedule(g *gocui.Gui) error {
	view, err := g.View(gui.views.schedule.name)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.schedule.name, err)
			return err
		}
	}
	view.Clear()
	gui.drawSchedule(view)
	gui.state.status = " press enter to save schedule "
	return nil
}

func (e *scheduleEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	if key == gocui.KeyArrowDown {
		_, cY := v.Cursor()
		if cY >= maxScheduleItems {
			return
		}
		v.MoveCursor(0, 1, false)
	}

	if key == gocui.KeyArrowUp {
		_, cY := v.Cursor()
		if cY <= 0 {
			return
		}
		v.MoveCursor(0, -1, false)
	}

	if key == gocui.KeyArrowRight {
		_, cY := v.Cursor()
		switch cY {
		case 0:
			if e.g.config.Schedule.Monthly >= maxSnapshotAmount {
				e.g.config.Schedule.Monthly = 99
				return
			}
			e.g.config.Schedule.Monthly++
		case 1:
			if e.g.config.Schedule.Weekly >= maxSnapshotAmount {
				e.g.config.Schedule.Weekly = 99
				return
			}
			e.g.config.Schedule.Weekly++
		case 2:
			if e.g.config.Schedule.Daily >= maxSnapshotAmount {
				e.g.config.Schedule.Daily = 99
				return
			}
			e.g.config.Schedule.Daily++
		case 3:
			if e.g.config.Schedule.Hourly >= maxSnapshotAmount {
				e.g.config.Schedule.Hourly = 99
				return
			}
			e.g.config.Schedule.Hourly++
		case 4:
			if e.g.config.Schedule.Boot >= maxSnapshotAmount {
				e.g.config.Schedule.Boot = 99
				return
			}
			e.g.config.Schedule.Boot++
		}
		err := e.g.updateSchedule(e.g.g)
		if err != nil {
			e.g.log.Errorf("can't update schedule: %v", err)
		}
	}

	if key == gocui.KeyArrowLeft {
		_, cY := v.Cursor()
		switch cY {
		case 0:
			if e.g.config.Schedule.Monthly <= 0 {
				e.g.config.Schedule.Monthly = 0
				return
			}
			e.g.config.Schedule.Monthly--
		case 1:
			if e.g.config.Schedule.Weekly <= 0 {
				e.g.config.Schedule.Weekly = 0
				return
			}
			e.g.config.Schedule.Weekly--
		case 2:
			if e.g.config.Schedule.Daily <= 0 {
				e.g.config.Schedule.Daily = 0
				return
			}
			e.g.config.Schedule.Daily--
		case 3:
			if e.g.config.Schedule.Hourly <= 0 {
				e.g.config.Schedule.Hourly = 0
				return
			}
			e.g.config.Schedule.Hourly--
		case 4:
			if e.g.config.Schedule.Boot <= 0 {
				e.g.config.Schedule.Boot = 0
				return
			}
			e.g.config.Schedule.Boot--
		}
		err := e.g.updateSchedule(e.g.g)
		if err != nil {
			e.g.log.Errorf("can't update schedule: %v", err)
		}
	}

}
