package gui

import (
	"errors"
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/jroimartin/gocui"
)

const maxScheduleItems = 4

type scheduleEditor struct {
	g *Gui
}

func (gui *Gui) scheduleView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.schedule.name,
		int(scheduleIndent*float32(gui.maxX)),
		gui.views.schedule.y0,
		gui.maxX,
		gui.maxY-volumesHeigh,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't set %s view: %v", gui.views.schedule.name, err)

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
		err = gui.g.SetKeybinding(gui.views.schedule.name, gocui.KeyEsc, gocui.ModNone, gui.escapeFromEditableView)
		if err != nil {
			return nil, err
		}
		err = gui.g.SetKeybinding(gui.views.schedule.name, gocui.KeyEnter, gocui.ModNone, gui.saveConfig)
		if err != nil {
			return nil, err
		}
	}

	return view, nil
}

func (gui *Gui) drawSchedule(view *gocui.View) {
	fmt.Fprintf(view, " Monthly:%2d \n", gui.config.Schedule.Monthly)
	fmt.Fprintf(view, " Weekly: %2d \n", gui.config.Schedule.Weekly)
	fmt.Fprintf(view, " Daily:  %2d \n", gui.config.Schedule.Daily)
	fmt.Fprintf(view, " Hourly: %2d \n", gui.config.Schedule.Hourly)
	fmt.Fprintf(view, " Boot:   %2d \n", gui.config.Schedule.Boot)
}

func (gui *Gui) editSchedule(g *gocui.Gui, view *gocui.View) error {
	err := gui.escapeFromViewsByName(gui.views.snapshots.name)
	if err != nil {
		return err
	}

	gui.g.Cursor = false
	view.Highlight = true
	view.Editable = true
	view.SelBgColor = gui.highlightBg

	_, err = gui.g.SetCurrentView(gui.views.schedule.name)
	if err != nil {
		return err
	}

	return err
}

func (gui *Gui) updateSchedule(g *gocui.Gui) error {
	view, err := g.View(gui.views.schedule.name)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't set %s view: %v", gui.views.schedule.name, err)

			return err
		}
	}
	view.Clear()
	gui.drawSchedule(view)
	gui.state.status = colors.FgRed("press enter to save schedule")

	return nil
}

func (e *scheduleEditor) Edit(view *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	if key == gocui.KeyArrowDown {
		_, cY := view.Cursor()
		if cY >= maxScheduleItems {
			return
		}
		view.MoveCursor(0, 1, false)
	}

	if key == gocui.KeyArrowUp {
		_, cY := view.Cursor()
		if cY <= 0 {
			return
		}
		view.MoveCursor(0, -1, false)
	}

	if key == gocui.KeyArrowRight {
		_, cY := view.Cursor()
		updated := e.g.config.Schedule.Increase(cY)
		if !updated {
			return
		}

		err := e.g.updateSchedule(e.g.g)
		if err != nil {
			logger.Client.ErrLog.Printf("can't update schedule: %v", err)
		}
	}

	if key == gocui.KeyArrowLeft {
		_, cY := view.Cursor()
		updated := e.g.config.Schedule.Decrease(cY)
		if !updated {
			return
		}

		err := e.g.updateSchedule(e.g.g)
		if err != nil {
			logger.Client.ErrLog.Printf("can't update schedule: %v", err)
		}
	}
}
