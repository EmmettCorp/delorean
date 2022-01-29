/*
Package domain keeps shared structures, constants and functions.
*/
package domain

import (
	"io/fs"
	"os"
)

// DeloreanPath is a path to delorean config and scripts.
const DeloreanPath = "/usr/local/delorean"

// DeloreanMountPoint is a mount point where subvolumes will be mount
// and where snapshots will be reachable.
const DeloreanMountPoint = "/run/delorean"

// Subvol5 represents subvolume with id = 5. A top level subvolume.
const Subvol5 = "subvol5"

// SnapshotsDirName represents main snapsshots root directory name.
const SnapshotsDirName = ".snapshots"

// Snapshots types.
const (
	Monthly = "monthly"
	Weekly  = "weekly"
	Daily   = "daily"
	Hourly  = "hourly"
	Boot    = "boot"
	Manual  = "manual"
	Restore = "restore"
)

const RWFileMode = 0o600

// CheckDir checks if directory exists and creates if it doesn't.
func CheckDir(ph string, fm fs.FileMode) error {
	_, err := os.Stat(ph)
	if os.IsNotExist(err) {
		return os.MkdirAll(ph, fm)
	}

	return err
}
