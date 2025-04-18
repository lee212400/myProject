package logger

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.StacktraceKey = "stacktrace"

	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.OutputPaths = []string{"stdout"}

	logger, err := cfg.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		panic(err)
	}

	log = logger.Sugar()
}

func Debug(msg string, args ...any) { log.Debugf(msg, args...) }
func Info(msg string, args ...any)  { log.Infof(msg, args...) }
func Warn(msg string, args ...any)  { log.Warnf(msg, args...) }
func Error(msg string, args ...any) { log.Errorf(msg, args...) }
