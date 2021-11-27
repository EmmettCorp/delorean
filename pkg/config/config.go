package config

// GuiConfig is for configuring visual things like colors and whether we show or
// hide things
type GuiConfig struct {
	// ScrollHeight determines how many characters you scroll at a time when
	// scrolling the main panel
	ScrollHeight int `yaml:"scrollHeight,omitempty"`

	// ScrollPastBottom determines whether you can scroll past the bottom of the
	// main view
	ScrollPastBottom bool `yaml:"scrollPastBottom,omitempty"`

	// IgnoreMouseEvents is for when you do not want to use your mouse to interact
	// with anything
	IgnoreMouseEvents bool `yaml:"mouseEvents,omitempty"`

	// Theme determines what colors and color attributes your panel borders have.
	// I always set inactiveBorderColor to black because in my terminal it's more
	// of a grey, but that doesn't work in your average terminal. I highly
	// recommended finding a combination that works for you
	Theme ThemeConfig `yaml:"theme,omitempty"`

	// ShowAllContainers determines whether the Containers panel contains all the
	// containers returned by `docker ps -a`, or just those containers that aren't
	// directly linked to a service. It is probably desirable to enable this if
	// you have multiple containers per service, but otherwise it can cause a lot
	// of clutter
	ShowAllContainers bool `yaml:"showAllContainers,omitempty"`

	// ReturnImmediately determines whether you get the 'press enter to return to
	// lazydocker' message after a subprocess has completed. You would set this to
	// true if you often want to see the output of subprocesses before returning
	// to lazydocker. I would default this to false but then people who want it
	// set to true won't even know the config option exists.
	ReturnImmediately bool `yaml:"returnImmediately,omitempty"`

	// WrapMainPanel determines whether we use word wrap on the main panel
	WrapMainPanel bool `yaml:"wrapMainPanel,omitempty"`
}

// UserConfig holds all of the user-configurable options
type UserConfig struct {
	// Gui is for configuring visual things like colors and whether we show or
	// hide things
	Gui GuiConfig `yaml:"gui,omitempty"`

	// OS determines what defaults are set for opening files and links
	OS OSConfig `yaml:"oS,omitempty"`

	// CommandTemplates determines what commands actually get called when we run
	// certain commands
	CommandTemplates CommandTemplatesConfig `yaml:"commandTemplates,omitempty"`
}

// ThemeConfig is for setting the colors of panels and some text.
type ThemeConfig struct {
	LightTheme                bool     `yaml:"lightTheme"`
	ActiveBorderColor         []string `yaml:"activeBorderColor"`
	InactiveBorderColor       []string `yaml:"inactiveBorderColor"`
	OptionsTextColor          []string `yaml:"optionsTextColor"`
	SelectedLineBgColor       []string `yaml:"selectedLineBgColor"`
	SelectedRangeBgColor      []string `yaml:"selectedRangeBgColor"`
	CherryPickedCommitBgColor []string `yaml:"cherryPickedCommitBgColor"`
	CherryPickedCommitFgColor []string `yaml:"cherryPickedCommitFgColor"`
}

// AppConfig contains the base configuration fields required for lazydocker.
type AppConfig struct {
	Name       string `long:"name" env:"NAME" default:"lazydocker"`
	Version    string `long:"version" env:"VERSION" default:"unversioned"`
	UserConfig *UserConfig
}

// NewAppConfig makes a new app config
func NewAppConfig(name, version string) (*AppConfig, error) {
	userConfig := UserConfig{}

	appConfig := &AppConfig{
		Name:       name,
		Version:    version,
		UserConfig: &userConfig,
	}

	return appConfig, nil
}

// OSConfig contains config on the level of the os
type OSConfig struct {
	// OpenCommand is the command for opening a file
	OpenCommand string `yaml:"openCommand,omitempty"`

	// OpenCommand is the command for opening a link
	OpenLinkCommand string `yaml:"openLinkCommand,omitempty"`
}

// CommandTemplatesConfig determines what commands actually get called when we
// run certain commands
type CommandTemplatesConfig struct {
	// Restore restores chosen snapshot.
	Restore string `yaml:"restore,omitempty"`

	BtrfsPath string `yaml:"btrfsPath,omitempty"`

	// CheckBtrfsPath
	CheckBtrfsPath string `yaml:"checkBtrfsPath,omitempty"`
}
