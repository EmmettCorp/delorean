package config

type (
	// UserConfig holds all of the user-configurable options
	UserConfig struct {
		// Gui is for configuring visual things like colors and whether we show or
		// hide things
		Gui GuiConfig `yaml:"gui,omitempty"`

		// OS determines what defaults are set for opening files and links
		OS OSConfig `yaml:"oS,omitempty"`

		// CommandTemplates determines what commands actually get called when we run
		// certain commands
		CommandTemplates CommandTemplatesConfig `yaml:"commandTemplates,omitempty"`

		// Refresh settings
		Refresher RefresherConfig `yaml:"refresher"`
	}

	// RefresherConfig keeps refres settings.
	RefresherConfig struct {
		RefreshInterval int `yaml:"refreshInterval"`
		FetchInterval   int `yaml:"fetchInterval"`
	}
)

func GetDefaultConfig() *UserConfig {
	return &UserConfig{
		Gui: GuiConfig{
			ScrollHeight:     2,
			ScrollPastBottom: true,
			MouseEvents:      true,
			Theme: ThemeConfig{
				LightTheme:                false,
				ActiveBorderColor:         []string{"green", "bold"},
				InactiveBorderColor:       []string{"white"},
				OptionsTextColor:          []string{"blue"},
				SelectedLineBgColor:       []string{"default"},
				SelectedRangeBgColor:      []string{"blue"},
				CherryPickedCommitBgColor: []string{"blue"},
				CherryPickedCommitFgColor: []string{"cyan"},
			},
		},
		Refresher: RefresherConfig{
			RefreshInterval: 10,
			FetchInterval:   60,
		},
	}
}
