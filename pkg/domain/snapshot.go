package domain

import (
	"fmt"
	"strings"
	"time"
)

// SnapshotFormat is a format of snapshots names.
const SnapshotFormat = "2006-01-02_15:04:05"

// Snapshot represents snapshot object, keeps all needed data.
type Snapshot struct {
	ID          string `json:"id"`
	Path        string `json:"path"`
	Label       string `json:"label"`
	Type        string `json:"type"` // manual, weekly, daily, etc.
	VolumeLabel string `json:"volume_label"`
	VolumeID    string `json:"volume_id"`
	Timestamp   int64  `json:"timestamp"`
}

// NewSnapshot creates a new snapshot object by path to snapshot, volume label and volume id.
// It is supposed that path to snapshots looks like `**/<volume>/<snapshot_type>/<snapshot_id>`.
// Example:
// 			/run/delorean/.snapshots/@/manual/2022-02-16_16:17:45
//
func NewSnapshot(ph, vLabel, vID string) (Snapshot, error) {
	sn := Snapshot{
		Path:        ph,
		VolumeLabel: vLabel,
		VolumeID:    vID,
	}

	ss := strings.Split(ph, "/")
	if len(ss) < 4 { // nolint:gomnd // there MUST be snapshots `type` and `id` in path
		return Snapshot{}, fmt.Errorf("path is too short `%s`", ph)
	}
	sn.Label = ss[len(ss)-1]
	sn.Type = ss[len(ss)-2]

	t, err := time.Parse(SnapshotFormat, sn.Label)
	if err != nil {
		return Snapshot{}, fmt.Errorf("can't parse label with snapshot format: %v", err)
	}
	sn.Timestamp = t.Unix()

	return sn, nil
}
