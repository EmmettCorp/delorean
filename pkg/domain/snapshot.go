package domain

import "strings"

type Snapshot struct {
	Path   string
	Label  string
	Type   string // manual, weekly, daily, etc.
	Volume string // volume label
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
