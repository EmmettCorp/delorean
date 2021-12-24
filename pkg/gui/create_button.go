package gui

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/colors"
	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/jroimartin/gocui"
)

func (gui *Gui) createButton() (*gocui.View, error) {
	view, err := gui.g.SetView(gui.views.createBtn.name,
		gui.views.createBtn.x0,
		gui.views.createBtn.y0,
		gui.views.createBtn.x1,
		gui.views.createBtn.y1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			gui.log.Errorf("can't build %s button: %v", gui.views.createBtn.name, err)
			return nil, err
		}
		err := gui.g.SetKeybinding(gui.views.createBtn.name, gocui.MouseLeft, gocui.ModNone, gui.createSnapshot)
		if err != nil {
			return nil, err
		}
		fmt.Fprint(view, gui.views.createBtn.name)
	}

	return view, nil
}

func (gui *Gui) createSnapshot(g *gocui.Gui, view *gocui.View) error {
	for _, vol := range gui.config.Volumes {
		if !vol.Active {
			continue
		}

		err := commands.CreateSnapshot(vol.Point, fmt.Sprintf("%s/%s", gui.config.SnapshotsPath, domain.Manual))
		if err != nil {
			return fmt.Errorf("can't create snapshot for %s: %v", vol.Point, err)
		}
	}
	gui.state.status = colors.FgGreen(" new snapshot is created ")

	return nil
}
