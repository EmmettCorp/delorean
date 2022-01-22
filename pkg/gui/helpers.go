package gui

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

func (gui *Gui) getVolumeByID(id string) (domain.Volume, error) {
	for i := range gui.config.Volumes {
		if gui.config.Volumes[i].ID == id {
			return gui.config.Volumes[i], nil
		}
	}

	return domain.Volume{}, fmt.Errorf("can't find volume by id `%s`", id)
}

func (gui *Gui) volumeInRootFs(vol domain.Volume) bool {
	return vol.Device.Path == gui.config.RootDevice
}
