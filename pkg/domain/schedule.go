package domain

const (
	minSnapshotAmount = 0
	maxSnapshotAmount = 99
)

const (
	monthlyType = iota
	weeklyType
	dailyType
	hourlyType
	bootType
)

// Schedule represents the settings of scheduled snapshots.
type Schedule struct {
	Monthly int `json:"monthly"`
	Weekly  int `json:"weekly"`
	Daily   int `json:"daily"`
	Hourly  int `json:"hourly"`
	Boot    int `json:"boot"`
}

// Increase increases value of `schedType` for settings.
func (s *Schedule) Increase(schedType int) bool {
	updated := false
	switch schedType {
	case monthlyType:
		if s.Monthly < maxSnapshotAmount {
			s.Monthly++
			updated = true
		}
	case weeklyType:
		if s.Weekly < maxSnapshotAmount {
			s.Weekly++
			updated = true
		}
	case dailyType:
		if s.Daily < maxSnapshotAmount {
			s.Daily++
			updated = true
		}
	case hourlyType:
		if s.Hourly < maxSnapshotAmount {
			s.Hourly++
			updated = true
		}
	case bootType:
		if s.Boot < maxSnapshotAmount {
			s.Boot++
			updated = true
		}
	}

	return updated
}

// Decrease decreases value of `schedType` for settings.
func (s *Schedule) Decrease(schedType int) bool {
	updated := false
	switch schedType {
	case monthlyType:
		if s.Monthly > minSnapshotAmount {
			s.Monthly--
			updated = true
		}
	case weeklyType:
		if s.Weekly > minSnapshotAmount {
			s.Weekly--
			updated = true
		}
	case dailyType:
		if s.Daily > minSnapshotAmount {
			s.Daily--
			updated = true
		}
	case hourlyType:
		if s.Hourly > minSnapshotAmount {
			s.Hourly--
			updated = true
		}
	case bootType:
		if s.Boot > minSnapshotAmount {
			s.Boot--
			updated = true
		}
	}

	return updated
}
