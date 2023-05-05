package log

import (
	"context"

	"go.uber.org/zap"
)

type ILogger interface {
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

var logger ILogger = newLogger()

func Default() ILogger {
	return logger
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	Default().Debug(ctx, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	Default().Info(ctx, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	Default().Warn(ctx, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	Default().Error(ctx, msg, fields...)
}
