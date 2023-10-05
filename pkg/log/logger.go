package log

import (
	"context"
	"time"

	"github.com/anguloc/zet/pkg/safe"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	handler *zap.Logger
	level   zapcore.Level
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...Field) {
	if l.level > zap.DebugLevel {
		return
	}
	l.handler.Debug(msg, fields...)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...Field) {
	if l.level > zap.InfoLevel {
		return
	}
	l.handler.Info(msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...Field) {
	if l.level > zap.WarnLevel {
		return
	}
	l.handler.Warn(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...Field) {
	if l.level > zap.ErrorLevel {
		return
	}
	l.handler.Error(msg, fields...)
}

func newLogger() *Logger {
	var cores []zapcore.Core

	cfg := zap.NewProductionConfig()
	ec := cfg.EncoderConfig
	ec.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	ec.TimeKey = "datetime"
	ec.CallerKey = "file"

	// >= warn
	cores = append(cores, zapcore.NewCore(
		zapcore.NewJSONEncoder(ec),
		zapcore.AddSync(zapcore.AddSync(&lumberjack.Logger{
			Filename:   safe.Path("log/error.log"),
			MaxSize:    512,
			MaxBackups: 3,
			LocalTime:  true,
		})),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zapcore.WarnLevel
		}),
	))

	// < warn
	cores = append(cores, zapcore.NewCore(
		zapcore.NewJSONEncoder(ec),
		zapcore.AddSync(zapcore.AddSync(&lumberjack.Logger{
			Filename:   safe.Path("log/access.log"),
			MaxSize:    512,
			MaxBackups: 3,
			LocalTime:  true,
		})),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level < zapcore.WarnLevel
		}),
	))

	var opt []zap.Option
	opt = append(opt, zap.AddCaller(), zap.AddCallerSkip(2))

	zapLogger := zap.New(zapcore.NewTee(cores...), opt...)
	return &Logger{
		handler: zapLogger,
		level:   zap.DebugLevel,
	}
}
