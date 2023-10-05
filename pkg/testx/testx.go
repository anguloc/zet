package testx

import (
	"context"

	"github.com/anguloc/zet/pkg/log"
	gLogger "gorm.io/gorm/logger"
)

func Context(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}

	// 单测打印sql
	ctx = log.WithGormContext(ctx, &log.LoggerGormVal{
		Level: gLogger.Info,
	})

	return ctx
}
