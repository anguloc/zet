package cron

import (
	"context"

	"github.com/anguloc/zet/pkg/log"
	"github.com/robfig/cron/v3"
)

type job struct {
	ctx      context.Context //上下文，执行时注入，由于cron组件没有ctx参数，这里转一层
	schedule cron.Schedule
	fn       IJob
}

func (j *job) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Error(j.ctx, "cronJobPanic", log.Any("panic", r))
			return
		}
	}()
	_ = j.fn.Run(j.ctx)
}

type IJob interface {
	Run(ctx context.Context) error
}

type jobFunc func(ctx context.Context) error

func (f jobFunc) Run(ctx context.Context) error {
	return f(ctx)
}
