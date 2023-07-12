package cron

import (
	"context"
	"time"

	"github.com/anguloc/zet/pkg/log"
)

type Worker struct {
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) Init(ctx context.Context) error {
	return nil
}

func (w *Worker) Run(ctx context.Context) error {
	if len(defaultCron.jobs) <= 0 {
		return nil
	}

	for _, j := range defaultCron.jobs {
		j.ctx = ctx
		defaultCron.cron.Schedule(j.schedule, j)
	}

	defaultCron.cron.Start()
	<-ctx.Done()
	log.Info(ctx, "cronReceiveEndSig")
	newCtx := defaultCron.cron.Stop()
	select {
	case <-newCtx.Done():
		log.Info(ctx, "cronEndSuccess")
	case <-time.After(time.Second * 5):
		log.Error(ctx, "cronEndFailed")
	}
	return nil
}