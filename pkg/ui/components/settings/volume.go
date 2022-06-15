package settings

import "github.com/EmmettCorp/delorean/pkg/domain"

type (
	volume struct {
		ID               string
		Subvol           string
		Label            string
		SnapshotsPath    string
		Active           bool
		DeviceUUID       string
		DevicePath       string
		DeviceMountPoint string
	}
)

// FilterValue is used to set filter item and required for `list.Model` interface.
func (v *volume) FilterValue() string { return v.Label }

func domainVolumeToVolume(dv domain.Volume) volume {
	return volume{
		ID:               dv.ID,
		Subvol:           dv.Subvol,
		Label:            dv.Label,
		SnapshotsPath:    dv.SnapshotsPath,
		Active:           dv.Active,
		DeviceUUID:       dv.Device.UUID,
		DevicePath:       dv.Device.Path,
		DeviceMountPoint: dv.Device.MountPoint,
	}
}
