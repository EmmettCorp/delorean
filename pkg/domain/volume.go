package domain

const (
	Monthly = "monthly"
	Weekly  = "weekly"
	Daily   = "daily"
	Hourly  = "hourly"
	Boot    = "boot"
	Manual  = "manual"
	Revert  = "revert"
)

// Volume represents btrfs volume.
type Volume struct {
	Subvol     string `json:"subvol"`
	Label      string `json:"label"`
	Device     string `json:"device"`
	UUID       string `json:"uuid"`
	MountPoint string `json:"mount_point"`
	Active     bool   `json:"active"`
	Mounted    bool   `json:"mounted"`
}
