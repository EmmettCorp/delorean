package domain

// Volume represents btrfs volume.
type Volume struct {
	Label  string `json:"label"`
	UID    string `json:"uid"`
	Active bool   `json:"active"`
}
