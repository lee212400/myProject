package logger

import (
	"github.com/lee212400/myProject/domain/entity"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.LevelKey = "logLevel"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // file形式(fullpathか該当ファイルのみか)
	cfg.Development = true                                      // 開発者モードでstack追跡範囲が広くなる

	base, err := cfg.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}

	log = base.Sugar()
}

func WithContext(ctx *entity.Context) *zap.SugaredLogger {

	fields := []zap.Field{}

	if rid, ok := ctx.Value("request-id").(string); ok {
		fields = append(fields, zap.String("request-id", rid))
	}
	if tid, ok := ctx.Value("trace-id").(string); ok {
		fields = append(fields, zap.String("trace-id", tid))
	}

	return log.Desugar().With(fields...).Sugar()
}

func Debug(msg string, args ...any) { log.Debugf(msg, args...) }
func Info(msg string, args ...any)  { log.Infof(msg, args...) }
func Warn(msg string, args ...any)  { log.Warnf(msg, args...) }
func Error(msg string, args ...any) { log.Errorf(msg, args...) }
