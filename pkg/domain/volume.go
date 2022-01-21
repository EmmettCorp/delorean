package domain

// Volume represents btrfs volume.
type Volume struct {
	ID            string `json:"id"`
	Subvol        string `json:"subvol"`
	Label         string `json:"label"`
	Device        string `json:"device"`
	UUID          string `json:"device_uuid"`
	MountPoint    string `json:"mount_point"`
	SnapshotsPath string `json:"snapshots_path"`
	Active        bool   `json:"active"`
	Mounted       bool   `json:"mounted"`
}

type Device struct {
	UUID       string `json:"uuid"`
	Path       string `json:"path"`
	MountPoint string `json:"mount_point"`
	Mounted    bool   `json:"mounted"`
}
