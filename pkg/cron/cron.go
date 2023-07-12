package cron

import (
	"context"
	"sync"
	"time"

	"github.com/anguloc/zet/pkg/log"
	"github.com/robfig/cron/v3"
)

type ICron interface {
	RegisterFunc(ctx context.Context, spec string, fn func(ctx context.Context) error) error
	RegisterJob(ctx context.Context, spec string, ij IJob) error
	Start(ctx context.Context) error
}

type Cron struct {
	mu   *sync.Mutex
	jobs []*job
}

func NewCron() *Cron {
	return &Cron{
		mu: &sync.Mutex{},
	}
}

var defaultCron = newCron()

func newCron() *cron.Cron {
	return cron.New(cron.WithLogger(cron.DiscardLogger))
}

func (c *Cron) RegisterFunc(ctx context.Context, spec string, fn func(ctx context.Context) error) error {
	return c.RegisterJob(ctx, spec, jobFunc(fn))
}

func (c *Cron) RegisterJob(ctx context.Context, spec string, ij IJob) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	sch, err := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).Parse(spec)
	if err != nil {
		return err
	}
	c.jobs = append(c.jobs, &job{
		schedule: sch,
		fn:       ij,
	})
	return nil
}

func (c *Cron) Start(ctx context.Context) error {
	if len(c.jobs) <= 0 {
		return nil
	}

	for _, j := range c.jobs {
		j.ctx = ctx
		defaultCron.Schedule(j.schedule, j)
	}

	defaultCron.Start()
	<-ctx.Done()
	log.Info(ctx, "cronReceiveEndSig")
	newCtx := defaultCron.Stop()
	select {
	case <-newCtx.Done():
		log.Info(ctx, "cronEndSuccess")
	case <-time.After(time.Second * 5):
		log.Error(ctx, "cronEndFailed")
	}
	return nil
}
