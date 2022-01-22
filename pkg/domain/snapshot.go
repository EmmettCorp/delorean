package domain

import (
	"strings"
	"time"
)

const SnapshotFormat = "2006-01-02_15:04:05"

type Snapshot struct {
	ID          int64
	Path        string
	Label       string
	Type        string // manual, weekly, daily, etc.
	VolumeLabel string
	VolumeID    string
	Timestamp   int64
}

func (s *Snapshot) SetLabel() {
	ss := strings.Split(s.Path, "/")
	if len(ss) == 0 {
		return
	}

	s.Label = ss[len(ss)-1]
}

func (s *Snapshot) SetType() {
	ss := strings.Split(s.Path, "/")
	if len(ss) < 2 {
		return
	}

	s.Type = ss[len(ss)-2]
}

func (s *Snapshot) SetTimestamp() {
	t, err := time.Parse(SnapshotFormat, s.Label)
	if err != nil {
		return
	}
	s.Timestamp = t.Unix()
}
