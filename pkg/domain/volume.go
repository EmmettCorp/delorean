package domain

// Volume represents btrfs volume.
type Volume struct {
	ID            string `json:"id"`
	Subvol        string `json:"subvol"`
	Label         string `json:"label"`
	Device        string `json:"device"`
	UUID          string `json:"uuid"`
	MountPoint    string `json:"mount_point"`
	SnapshotsPath string `json:"snapshots_path"`
	Active        bool   `json:"active"`
	Mounted       bool   `json:"mounted"`
}
