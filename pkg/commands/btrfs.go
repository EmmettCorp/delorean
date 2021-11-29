package commands

import (
	"sync"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/i18n"
	"github.com/imdario/mergo"
	"github.com/sirupsen/logrus"
)

// BtrfsCommand is our main docker interface
type BtrfsCommand struct {
	Log                    *logrus.Entry
	OSCommand              *OSCommand
	Tr                     *i18n.TranslationSet
	Config                 *config.AppConfig
	InDockerComposeProject bool
	ShowExited             bool
	ErrorChan              chan error
	ContainerMutex         sync.Mutex
	ServiceMutex           sync.Mutex
}

// NewBtrfsCommand it runs docker commands
func NewBtrfsCommand(log *logrus.Entry, osCommand *OSCommand, tr *i18n.TranslationSet, config *config.AppConfig, errorChan chan error) (*BtrfsCommand, error) {

	btrfsCommand := &BtrfsCommand{
		Log:                    log,
		OSCommand:              osCommand,
		Tr:                     tr,
		Config:                 config,
		ErrorChan:              errorChan,
		ShowExited:             true,
		InDockerComposeProject: true,
	}

	// command := utils.ApplyTemplate(
	// 	config.UserConfig.CommandTemplates.CheckBtrfsPath,
	// 	btrfsCommand.NewCommandObject(CommandObject{}),
	// )

	// log.Warn(command)

	// err := osCommand.RunCommand(
	// 	utils.ApplyTemplate(
	// 		config.UserConfig.CommandTemplates.CheckBtrfsPath,
	// 		btrfsCommand.NewCommandObject(CommandObject{}),
	// 	),
	// )
	// if err != nil {
	// 	btrfsCommand.InDockerComposeProject = false
	// 	log.Warn(err.Error())
	// }

	return btrfsCommand, nil
}

// CommandObject is what we pass to our template resolvers when we are running a custom command. We do not guarantee that all fields will be populated: just the ones that make sense for the current context
type CommandObject struct {
	BtrfsPath string
}

// NewCommandObject takes a command object and returns a default command object with the passed command object merged in
func (c *BtrfsCommand) NewCommandObject(obj CommandObject) CommandObject {
	defaultObj := CommandObject{BtrfsPath: c.Config.UserConfig.CommandTemplates.BtrfsPath}
	mergo.Merge(&defaultObj, obj)
	return defaultObj
}
