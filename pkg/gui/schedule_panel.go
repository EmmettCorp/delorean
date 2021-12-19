package gui

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) scheduleView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.schedule.name,
		int(0.8*float32(gui.maxX)),
		gui.views.schedule.y0,
		gui.maxX,
		gui.maxY-5,
	)
	view.Editable = true
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			gui.log.Errorf("can't set %s view: %v", gui.views.schedule.name, err)
			return nil, err
		}

		view.Title = gui.views.schedule.name
		gui.drawSchedule(view)
	}

	return view, nil
}

func (gui *Gui) drawSchedule(view *gocui.View) {
	fmt.Fprintf(view, " Monthly: %d\n", gui.config.Schedule.Monthly)
	fmt.Fprintf(view, " Weekly:  %d\n", gui.config.Schedule.Weekly)
	fmt.Fprintf(view, " Daily:   %d\n", gui.config.Schedule.Daily)
	fmt.Fprintf(view, " Hourly:  %d\n", gui.config.Schedule.Hourly)
	fmt.Fprintf(view, " Boot:    %d\n", gui.config.Schedule.Boot)
}
