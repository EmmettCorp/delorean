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

// Scripts and systemd units path.
const (
	ScriptsDirectoryName = "scripts"
	SystemdDirectoryName = "systemd"
)

type SnapshotType string

// Snapshots types.
const (
	Monthly SnapshotType = "monthly"
	Weekly  SnapshotType = "weekly"
	Daily   SnapshotType = "daily"
	Hourly  SnapshotType = "hourly"
	Boot    SnapshotType = "boot"
	Manual  SnapshotType = "manual"
	Restore SnapshotType = "restore"
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

func (s SnapshotType) String() string {
	return string(s)
}
