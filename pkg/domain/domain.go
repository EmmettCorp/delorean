/*
Package domain keeps shared structures, constants and functions.
*/
package domain

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
	Revert  = "revert"
)
