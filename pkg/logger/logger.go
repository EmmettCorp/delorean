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

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	defaultLogDir = "/var/log/delorean"
)

type Client struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

func New() (*Client, error) {
	err := checkDir(defaultLogDir, domain.RWFileMode)
	if err != nil {
		return nil, err
	}

	ph := path.Join(defaultLogDir, "app.log")

	logFile, err := os.OpenFile(ph, os.O_RDWR|os.O_CREATE|os.O_TRUNC, domain.RWFileMode) // nolint gosec: ph is constructed from constants.
	if err != nil {
		return nil, err
	}

	infoLog := log.New(logFile, "info: ", log.Ltime|log.Lshortfile)
	errLog := log.New(logFile, "err: ", log.Ltime|log.Lshortfile)

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
