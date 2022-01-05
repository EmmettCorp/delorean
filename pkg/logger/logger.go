package logger

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultLogDir = "/var/log/delorean"
	logNameFormat = "2006-01-02_15-04-05"
)

type Client struct {
	*zap.SugaredLogger
}

func New() (*Client, error) {
	err := checkDir(defaultLogDir, 0600)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s.log", defaultLogDir, time.Now().Format(logNameFormat))

	zc := zap.NewProductionConfig()
	zc.OutputPaths = []string{path}
	zc.ErrorOutputPaths = []string{path}
	zc.DisableCaller = true
	zc.DisableStacktrace = true
	zc.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.RFC3339))
	}

	zl, err := zc.Build([]zap.Option{}...)
	if err != nil {
		return nil, err
	}

	defer zl.Sync()

	sugar := zl.Sugar()

	return &Client{sugar}, nil
}

// CloseOrLog is a helper for any defer closer.Close() call.
func (lc *Client) CloseOrLog(c io.Closer) {
	err := c.Close()
	if err != nil {
		lc.Errorf("fail to close: %v", err)
	}
}

func checkDir(path string, fileMode fs.FileMode) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Mkdir(path, fileMode)
	}

	return err
}
