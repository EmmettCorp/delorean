/*
Package logger is responsible for logger initialization.
*/
package logger

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"time"
)

const (
	defaultLogDir = "/var/log/delorean"
	logNameFormat = "2006-01-02_15:04:05"
)

type Client struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

func New() (*Client, error) {
	err := checkDir(defaultLogDir, 0o600)
	if err != nil {
		return nil, err
	}

	ph := path.Join(defaultLogDir, fmt.Sprintf("%s.log", time.Now().Format(logNameFormat)))

	logFile, err := os.OpenFile(ph, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600) // nolint gosec: ph is constructed from constants.
	if err != nil {
		return nil, err
	}

	infoLog := log.New(logFile, "info:\t", log.Ltime)
	errLog := log.New(logFile, "err:\t", log.Ltime)

	return &Client{
		InfoLog: infoLog,
		ErrLog:  errLog,
	}, nil
}

// CloseOrLog is a helper for any defer closer.Close() call.
func (lc *Client) CloseOrLog(c io.Closer) {
	err := c.Close()
	if err != nil {
		lc.ErrLog.Printf("fail to close: %v", err)
	}
}

func checkDir(ph string, fileMode fs.FileMode) error {
	_, err := os.Stat(ph)
	if os.IsNotExist(err) {
		return os.Mkdir(ph, fileMode)
	}

	return err
}

// Errorf logs error with message from `format`.
// `format` could be just a message for error.
// `params` are parameters for format.
func (lc *Client) Errorf(err error, format string, params ...interface{}) {
	format = fmt.Sprintf(format, params...)
	if format == "" {
		lc.ErrLog.Printf("%v", err)

		return
	}
	lc.ErrLog.Printf("%s: %v", format, err)
}

// Error logs error with message.
func (lc *Client) Error(message string) {
	lc.ErrLog.Print(message)
}

// Infof logs info message from `format` with `params`.
func (lc *Client) Infof(format string, params ...interface{}) {
	lc.InfoLog.Printf(format, params...)
}

// Info logs info `message`.
func (lc *Client) Info(message string) {
	lc.InfoLog.Print(message)
}
