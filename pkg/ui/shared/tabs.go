package shared

type TabItem int

const (
	SnapshotsTab TabItem = iota
	ScheduleTab
	SettingsTab
	AnyTab // used for elements that don't belong to any tab. As example, tabs themselves.
)

const (
	snapshotsTabTitle = "Snapshots"
	scheduleTabTitle  = "Schedule"
	settingsTabTitle  = "Settings"
)

func (t TabItem) String() string {
	switch t {
	case SnapshotsTab:
		return snapshotsTabTitle
	case ScheduleTab:
		return scheduleTabTitle
	case SettingsTab:
		return settingsTabTitle
	default:
		return ""
	}
}

// GetTabItems returns the list of all available (visually) tabs.
func GetTabItems() []TabItem {
	return []TabItem{
		SnapshotsTab,
		SettingsTab,
	}
}
