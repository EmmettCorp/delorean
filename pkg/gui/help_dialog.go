package gui

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/jroimartin/gocui"
)

const (
	helpViewWidth = 55
	helpViewHeigh = 20
	two           = 2
	contentLength = helpViewWidth - 5
)

func (gui *Gui) helpView() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.helpView.name,
		gui.maxX/two-helpViewWidth/two,
		gui.maxY/two-helpViewHeigh/two,
		gui.maxX/two+helpViewWidth/two,
		gui.maxY/two+helpViewHeigh/two+helpViewHeigh%two,
	)
	if err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			logger.Client.ErrLog.Printf("can't set %s view: %v", gui.views.helpView.name, err)

			return nil, err
		}

		view.Editable = false
		view.Frame = true
		view.Wrap = false
		gui.g.Cursor = false
		view.Autoscroll = false

		content := gui.buildHelpContent()
		gui.views.helpView.lines = bytes.Count(content, []byte{'\n'})
		fmt.Fprintf(view, "%s\n", content)

		err := gui.setKeybindings(gui.getHelpKeybindings())
		if err != nil {
			return nil, err
		}
	}
	gui.views.helpView.visible = true

	return view, nil
}

func (gui *Gui) deleteHelpView() error {
	_, err := gui.g.View(gui.views.helpView.name)
	if err != nil {
		return nil // nolint
	}
	gui.views.helpView.visible = false

	return gui.g.DeleteView(gui.views.helpView.name)
}

func (gui *Gui) buildHelpContent() []byte {
	var content []byte

	divider := []byte(strings.Repeat("-", helpViewWidth-two))
	divider = append(divider, '\n')

	content = append(content, []byte(" General keybindings\n")...)
	content = append(content, divider...)
	content = append(content, gui.generalKeybindingsInfo()...)
	content = append(content, '\n')

	content = append(content, []byte(" Schedule keybindings\n")...)
	content = append(content, divider...)
	content = append(content, gui.scheduleKeybindingsInfo()...)
	content = append(content, '\n')

	content = append(content, []byte(" Snapshots keybindings\n")...)
	content = append(content, divider...)
	content = append(content, gui.snapshotsKeybindingsInfo()...)
	content = append(content, '\n')

	content = append(content, []byte(" Volumes keybindings\n")...)
	content = append(content, divider...)
	content = append(content, gui.volumesKeybindingsInfo()...)
	content = append(content, '\n')

	return content
}

func getKeybindingDescription(kb *binding) string {
	return fmt.Sprintf(" %s %s\n",
		colors.Paint(kb.Name, colors.Yellow),
		fmt.Sprintf(
			fmt.Sprintf("%%%ds", contentLength-len(kb.Name)), kb.Description),
	)
}

func (gui *Gui) getHelpKeybindings() []*binding {
	return []*binding{
		{
			ViewName: gui.views.helpView.name,
			Key:      gocui.MouseWheelUp,
			Modifier: gocui.ModNone,
			Handler:  scrollDown,
		},
		{
			ViewName: gui.views.helpView.name,
			Key:      gocui.MouseWheelDown,
			Modifier: gocui.ModNone,
			Handler:  gui.helpScrollUp,
		},
	}
}

func (gui *Gui) helpScrollUp(g *gocui.Gui, view *gocui.View) error {
	ox, oy := view.Origin()
	if oy >= gui.views.helpView.lines-helpViewHeigh {
		return nil
	}

	return view.SetOrigin(ox, oy+1)
}

func (gui *Gui) generalKeybindingsInfo() []byte {
	content := []byte{}
	for _, kb := range gui.getGeneralKeybindings() {
		if kb.Name == "" {
			continue
		}
		content = append(content, []byte(getKeybindingDescription(kb))...)
	}

	return content
}

func (gui *Gui) scheduleKeybindingsInfo() []byte {
	content := []byte{}
	schedKb := gui.getScheduleKeybindings()
	schedKb = append(schedKb, []*binding{
		{
			Name:        "Arrow left",
			Description: "Decrease items value",
		},
		{
			Name:        "Arrow right",
			Description: "Increase items value",
		},
		{
			Name:        "Arrow up",
			Description: "Move cursor up",
		},
		{
			Name:        "Arrow down",
			Description: "Move cursor down",
		},
	}...)
	for _, kb := range schedKb {
		if kb.Name == "" {
			continue
		}
		content = append(content, []byte(getKeybindingDescription(kb))...)
	}

	return content
}

func (gui *Gui) snapshotsKeybindingsInfo() []byte {
	content := []byte{}
	snapKb := []*binding{
		{
			Name:        "Arrow up",
			Description: "Move cursor up",
		},
		{
			Name:        "Arrow down",
			Description: "Move cursor down",
		},
	}
	for _, kb := range snapKb {
		if kb.Name == "" {
			continue
		}
		content = append(content, []byte(getKeybindingDescription(kb))...)
	}

	return content
}

func (gui *Gui) volumesKeybindingsInfo() []byte {
	content := []byte{}
	volKb := []*binding{
		{
			Name:        "Left mouse click",
			Description: "Toggle volumes observability",
		},
		{
			Name:        "Enter",
			Description: "Save volumes config",
		},
	}
	for _, kb := range volKb {
		if kb.Name == "" {
			continue
		}
		content = append(content, []byte(getKeybindingDescription(kb))...)
	}

	return content
}
