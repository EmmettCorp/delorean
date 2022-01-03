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
	Label  string `json:"label"`
	Point  string `json:"point"`
	Active bool   `json:"active"`
	Device string `json:"device"`
}
