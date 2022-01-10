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
	ID         string `json:"id"`
	Label      string `json:"label"`
	Device     string `json:"device"`
	UUID       string `json:"uuid"`
	MountPoint string `json:"mount_point"`
	Pluggable  bool   `json:"pluggable"`
	Active     bool   `json:"active"`
}
