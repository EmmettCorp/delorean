package shared

// TabItem is kind of identifier for tabs.
// It differs from `elements/tab` which is responsible for look and style, and which is data agnostic.
// TabItem, on the contrary, represents a concrete tab with a concrete title and related with a specific content.
type TabItem int

// TabItem tabs.
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

// String is a string representation of TabItems.
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
