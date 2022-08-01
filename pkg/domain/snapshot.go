package domain

import (
	"path"
	"time"
)

// SnapshotFormat is a format of snapshots names.
const SnapshotFormat = "2006-01-02_15:04:05"

// Snapshot represents snapshot object, keeps all needed data.
type Snapshot struct {
	Path      string       `json:"path"`
	Label     string       `json:"label"`
	Type      SnapshotType `json:"type"` // manual, weekly, daily, etc.
	Timestamp int64        `json:"timestamp"`
	Kernel    string       `json:"kernel"`
	Volume    Volume       `json:"volume"`
}

type SortableSnapshots []Snapshot

func (ss SortableSnapshots) Len() int           { return len(ss) }
func (ss SortableSnapshots) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }
func (ss SortableSnapshots) Less(i, j int) bool { return ss[i].Timestamp > ss[j].Timestamp }

// NewSnapshot creates a new snapshot object.
func NewSnapshot(phToSnapshots, kernel string, vol Volume, sType SnapshotType) Snapshot {
	ts := time.Now()
	label := ts.Format(SnapshotFormat)
	ph := path.Join(phToSnapshots, sType.String(), label)
	sn := Snapshot{
		Path:      ph,
		Label:     label,
		Type:      sType,
		Timestamp: ts.Unix(),
		Kernel:    kernel,
		Volume:    vol,
	}

	return sn
}
