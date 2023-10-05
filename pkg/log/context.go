package log

import (
	"context"

	gLogger "gorm.io/gorm/logger"
)

type loggerGormKey struct{}

type LoggerGormVal struct {
	Level gLogger.LogLevel
}

// FromGormContext ...
func FromGormContext(ctx context.Context) (*LoggerGormVal, bool) {
	if ctx == nil {
		return nil, false
	}
	l, ok := ctx.Value(loggerGormKey{}).(*LoggerGormVal)
	return l, ok
}

// WithGormContext ...
func WithGormContext(ctx context.Context, l *LoggerGormVal) context.Context {
	return context.WithValue(ctx, loggerGormKey{}, l)
}
