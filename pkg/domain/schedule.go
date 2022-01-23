package domain

const maxSnapshotAmount = 99

type Schedule struct {
	Monthly int `json:"monthly"`
	Weekly  int `json:"weekly"`
	Daily   int `json:"daily"`
	Hourly  int `json:"hourly"`
	Boot    int `json:"boot"`
}

func (s *Schedule) Increase(schedType int) bool {
	updated := false
	switch schedType {
	case 0:
		if s.Monthly < maxSnapshotAmount {
			s.Monthly++
			updated = true
		}
	case 1:
		if s.Weekly < maxSnapshotAmount {
			s.Weekly++
			updated = true
		}
	case 2:
		if s.Daily < maxSnapshotAmount {
			s.Daily++
			updated = true
		}
	case 3:
		if s.Hourly < maxSnapshotAmount {
			s.Hourly++
			updated = true
		}
	case 4:
		if s.Boot < maxSnapshotAmount {
			s.Boot++
			updated = true
		}
	}

	return updated
}

func (s *Schedule) Decrease(schedType int) bool {
	updated := false
	switch schedType {
	case 0:
		if s.Monthly > 0 {
			s.Monthly--
			updated = true
		}
	case 1:
		if s.Weekly > 0 {
			s.Weekly--
			updated = true
		}
	case 2:
		if s.Daily > 0 {
			s.Daily--
			updated = true
		}
	case 3:
		if s.Hourly > 0 {
			s.Hourly--
			updated = true
		}
	case 4:
		if s.Boot > 0 {
			s.Boot--
			updated = true
		}
	}

	return updated
}
