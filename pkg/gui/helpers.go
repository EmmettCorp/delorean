package gui

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

func (gui *Gui) getVolumeByUUID(uid string) (domain.Volume, error) {
	for i := range gui.config.Volumes {
		if gui.config.Volumes[i].Device.UUID == uid {
			return gui.config.Volumes[i], nil
		}
	}

	return domain.Volume{}, fmt.Errorf("can't find volume by uuid `%s`", uid)
}

func (gui *Gui) volumeInRootFs(vol domain.Volume) bool {
	return vol.Device.Path == gui.config.RootDevice
}
