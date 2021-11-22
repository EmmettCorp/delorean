package config

import (
	"os"
)

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
}

// ThemeConfig is for setting the colors of panels and some text.
type ThemeConfig struct {
	ActiveBorderColor   []string `yaml:"activeBorderColor,omitempty"`
	InactiveBorderColor []string `yaml:"inactiveBorderColor,omitempty"`
	OptionsTextColor    []string `yaml:"optionsTextColor,omitempty"`
}

// AppConfig contains the base configuration fields required for lazydocker.
type AppConfig struct {
	Debug       bool   `long:"debug" env:"DEBUG" default:"false"`
	Version     string `long:"version" env:"VERSION" default:"unversioned"`
	Commit      string `long:"commit" env:"COMMIT"`
	BuildDate   string `long:"build-date" env:"BUILD_DATE"`
	Name        string `long:"name" env:"NAME" default:"lazydocker"`
	BuildSource string `long:"build-source" env:"BUILD_SOURCE" default:""`
	UserConfig  *UserConfig
	ConfigDir   string
	ProjectDir  string
}

// NewAppConfig makes a new app config
func NewAppConfig(name, version, commit, date string, buildSource string, debuggingFlag bool, composeFiles []string, projectDir string) (*AppConfig, error) {
	userConfig := UserConfig{}

	appConfig := &AppConfig{
		Name:        name,
		Version:     version,
		Commit:      commit,
		BuildDate:   date,
		Debug:       debuggingFlag || os.Getenv("DEBUG") == "TRUE",
		BuildSource: buildSource,
		UserConfig:  &userConfig,
		ProjectDir:  projectDir,
	}

	return appConfig, nil
}
