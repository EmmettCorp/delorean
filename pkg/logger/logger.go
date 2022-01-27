/*
Package logger is responsible for logger initialization.
*/
package logger

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"sync"
)

const (
	defaultLogDir = "/var/log/delorean"
	RWFileMode    = 0o600 // duplicate domain but in order do not have conflicts
)

var once sync.Once // nolint gochecknoglobals: used only in this file

// Client is a singleton logger instance.
var Client *Instance // nolint gochecknoglobals: global on purpose

type Instance struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
	logFile *os.File
}

// Init creates a new singleton logger.
func Init() error {
	var err error
	once.Do(
		func() {
			Client, err = newInstance()
		})

	return err
}

func newInstance() (*Instance, error) {
	err := checkDir(defaultLogDir, RWFileMode)
	if err != nil {
		return nil, err
	}

	ph := path.Join(defaultLogDir, "app.log")

	logFile, err := os.OpenFile(ph, os.O_RDWR|os.O_CREATE|os.O_TRUNC, RWFileMode) // nolint gosec: ph is constructed from constants.
	if err != nil {
		return nil, err
	}

	infoLog := log.New(logFile, "info: ", log.Ltime|log.Lshortfile)
	errLog := log.New(logFile, "err: ", log.Ltime|log.Lshortfile)

	return &Instance{
		InfoLog: infoLog,
		ErrLog:  errLog,
		logFile: logFile,
	}, nil
}

// CloseOrLog is a helper for any defer closer.Close() call.
func (lc *Instance) CloseOrLog(c io.Closer) {
	err := c.Close()
	if err != nil {
		lc.ErrLog.Printf("fail to close: %v", err)
	}
}

// Close closes log file.
func (lc *Instance) Close() error {
	return lc.logFile.Close()
}

func checkDir(ph string, fileMode fs.FileMode) error {
	_, err := os.Stat(ph)
	if os.IsNotExist(err) {
		return os.Mkdir(ph, fileMode)
	}

	return err
}
