package shared

type TabItem int

const (
	SnapshotsTab TabItem = iota
	ScheduleTab
	SettingsTab
)

func (t TabItem) String() string {
	switch t {
	case SnapshotsTab:
		return "Snapshots"
	case ScheduleTab:
		return "Schedule"
	case SettingsTab:
		return "Settings"
	default:
		return ""
	}
}

func GetTabItems() []TabItem {
	return []TabItem{
		SnapshotsTab,
		SettingsTab,
	}
}
