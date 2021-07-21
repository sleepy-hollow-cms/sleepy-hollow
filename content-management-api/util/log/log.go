package log

import (
	"content-management-api/util/config"
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

	err := config.Config.Load()
	if err != nil {
		panic(err)
	}

	outputPaths := []string{}
	outputPaths = append(outputPaths, config.Config.Log.Output)
	var ll logLevel = logLevel(config.Config.Log.Level)

	logConfig := zap.Config{
		OutputPaths: outputPaths,
		Level:       zap.NewAtomicLevelAt(ll.zapLogLevel()),
		Encoding:    config.Config.Log.Encoding,
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