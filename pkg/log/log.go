package log

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/sirupsen/logrus"
)

// NewLogger returns a new logger
func NewLogger(config *config.AppConfig, rollrusHook string) *logrus.Entry {
	var log *logrus.Logger

	log = newProductionLogger()

	log.Formatter = &logrus.JSONFormatter{}

	return log.WithFields(logrus.Fields{
		"version": config.Version,
	})
}

func getLogLevel() logrus.Level {
	strLevel := os.Getenv("LOG_LEVEL")
	level, err := logrus.ParseLevel(strLevel)
	if err != nil {
		return logrus.DebugLevel
	}
	return level
}

func newDevelopmentLogger(config *config.AppConfig) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(getLogLevel())
	file, err := os.OpenFile("development.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("unable to log to file")
		os.Exit(1)
	}
	log.SetOutput(file)
	return log
}

func newProductionLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	log.SetLevel(logrus.ErrorLevel)
	return log
}
