package log

import (
	"context"
	"fmt"
	"time"

	"github.com/anguloc/zet/pkg/safe"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	gLogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	fields   []Field
	handler  *zap.Logger
	LogLevel gLogger.LogLevel
}

func NewGormLogger(ctx context.Context) *GormLogger {
	cfg := zap.NewProductionConfig()
	ec := cfg.EncoderConfig
	ec.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	ec.TimeKey = "datetime"
	ec.CallerKey = "file"

	var cores []zapcore.Core
	cores = append(cores, zapcore.NewCore(
		zapcore.NewJSONEncoder(ec),
		zapcore.AddSync(zapcore.AddSync(&lumberjack.Logger{
			Filename:   safe.Path("log/sql.log"),
			MaxSize:    512,
			MaxBackups: 3,
			LocalTime:  true,
		})),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			// 外层控制级别
			return true
		}),
	))

	var opt []zap.Option
	opt = append(opt, zap.AddCaller(), zap.AddCallerSkip(1))

	zapLogger := zap.New(zapcore.NewTee(cores...), opt...)

	level := gLogger.Warn
	if val, ok := FromGormContext(ctx); ok {
		level = val.Level
	}

	return &GormLogger{
		fields:   []Field{String("_cate_", "sql")},
		LogLevel: level,
		handler:  zapLogger,
	}
}

func (g *GormLogger) LogMode(level gLogger.LogLevel) gLogger.Interface {
	nl := *g
	nl.LogLevel = level
	return &nl
}

func (g *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel < gLogger.Info {
		return
	}
	g.handler.Info(fmt.Sprintf(s, i), g.fields...)
}

func (g *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel < gLogger.Warn {
		return
	}
	g.handler.Warn(fmt.Sprintf(s, i), g.fields...)
}

func (g *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.LogLevel < gLogger.Error {
		return
	}
	g.handler.Error(fmt.Sprintf(s, i), g.fields...)
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.LogLevel < gLogger.Info {
		return
	}
	sql, rows := fc()
	elapsed := time.Since(begin)
	// sql超过2000则截断
	if len(sql) > 2000 {
		sql = sql[:2000] + "..."
	}
	g.handler.Debug("sql stat", append(g.fields,
		Stringer("elapsed", elapsed),
		Int64("rows", rows),
		String("sql", sql),
	)...)
}
