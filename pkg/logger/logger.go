package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(path string) (*zap.SugaredLogger, error) {
	zc := zap.NewProductionConfig()
	zc.OutputPaths = []string{path}
	zc.DisableCaller = true
	zc.DisableStacktrace = true
	zc.ErrorOutputPaths = []string{path}
	zc.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.RFC3339))
	}

	zl, err := zc.Build([]zap.Option{}...)
	if err != nil {
		return nil, err
	}

	defer zl.Sync()

	sugar := zl.Sugar()

	return sugar, nil
}
