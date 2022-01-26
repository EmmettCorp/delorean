package domain

// Volume represents btrfs volume.
type Volume struct {
	ID            string `json:"id"`
	Subvol        string `json:"subvol"`
	Label         string `json:"label"`
	SnapshotsPath string `json:"snapshots_path"`
	Active        bool   `json:"active"`
	Device        Device `json:"device"`
}

// Device represent LVM device.
type Device struct {
	UUID       string `json:"uuid"`
	Path       string `json:"path"`
	MountPoint string `json:"mount_point"`
}
