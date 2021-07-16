package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logLevel string

func (ll logLevel) zapLogLevel() zapcore.Level {
	switch ll {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

type logger struct {
	*zap.SugaredLogger
}

func (l *logger) Error(err error) zap.Field {
	return zap.Error(err)
}

var (
	Logger *logger
)

func init() {
	outputPaths := []string{}
	outputPaths = append(outputPaths, "stdout")
	var ll logLevel = "debug"

	logConfig := zap.Config{
		OutputPaths: outputPaths,
		Level:       zap.NewAtomicLevelAt(ll.zapLogLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	l, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	Logger = &logger{l.Sugar()}
}
