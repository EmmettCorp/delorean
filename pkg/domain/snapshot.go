package domain

import (
	"fmt"
	"path"
	"strings"
	"time"
)

// SnapshotFormat is a format of snapshots names.
const SnapshotFormat = "2006-01-02_15:04:05"

// Snapshot represents snapshot object, keeps all needed data.
type Snapshot struct {
	Path        string `json:"path"`
	Label       string `json:"label"`
	Type        string `json:"type"` // manual, weekly, daily, etc.
	VolumeLabel string `json:"volume_label"`
	VolumeID    string `json:"volume_id"`
	Timestamp   int64  `json:"timestamp"`
	Kernel      string `json:"kernel"`
}

type SortableSnapshots []Snapshot

func (ss SortableSnapshots) Len() int           { return len(ss) }
func (ss SortableSnapshots) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }
func (ss SortableSnapshots) Less(i, j int) bool { return ss[i].Timestamp > ss[j].Timestamp }

// NewSnapshot creates a new snapshot object.
func NewSnapshot(phToSnapshots, sType, vLabel, vID, kernel string) Snapshot {
	ts := time.Now()
	label := ts.Format(SnapshotFormat)
	ph := path.Join(phToSnapshots, sType, label)
	sn := Snapshot{
		Path:        ph,
		VolumeLabel: vLabel,
		Label:       label,
		VolumeID:    vID,
		Type:        sType,
		Timestamp:   ts.Unix(),
		Kernel:      kernel,
	}

	return sn
}

// SnapshotByPath builds a snapshot object by path to snapshot, volume label and volume id.
// It is supposed that path to snapshots looks like `**/<volume>/<snapshot_type>/<snapshot_id>`.
// Example:
// 			/run/delorean/.snapshots/@/manual/2022-02-16_16:17:45
//
func SnapshotByPath(ph, vLabel, vID string) (Snapshot, error) {
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
