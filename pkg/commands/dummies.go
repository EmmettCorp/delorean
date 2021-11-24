package commands

import (
	"io/ioutil"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/i18n"
	"github.com/sirupsen/logrus"
)

// This file exports dummy constructors for use by tests in other packages

// NewDummyOSCommand creates a new dummy OSCommand for testing
func NewDummyOSCommand() *OSCommand {
	return NewOSCommand(NewDummyLog(), NewDummyAppConfig())
}

// NewDummyAppConfig creates a new dummy AppConfig for testing
func NewDummyAppConfig() *config.AppConfig {
	appConfig := &config.AppConfig{
		Name:    "test",
		Version: "unversioned",
	}
	return appConfig
}

// NewDummyLog creates a new dummy Log for testing
func NewDummyLog() *logrus.Entry {
	log := logrus.New()
	log.Out = ioutil.Discard
	return log.WithField("test", "test")
}

// NewDummyBtrfsCommand creates a new dummy BtrfsCommand for testing
func NewDummyBtrfsCommand() *BtrfsCommand {
	return NewDummyBtrfsCommandWithOSCommand(NewDummyOSCommand())
}

// NewDummyBtrfsCommandWithOSCommand creates a new dummy BtrfsCommand for testing
func NewDummyBtrfsCommandWithOSCommand(osCommand *OSCommand) *BtrfsCommand {
	return &BtrfsCommand{
		Log:       NewDummyLog(),
		OSCommand: osCommand,
		Tr:        i18n.NewTranslationSet(NewDummyLog()),
		Config:    NewDummyAppConfig(),
	}
}
