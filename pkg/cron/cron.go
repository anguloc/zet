package cron

import (
	"context"
	"sync"

	"github.com/robfig/cron/v3"
)

type ICron interface {
	RegisterFunc(spec string, fn func(ctx context.Context) error) error
	RegisterJob(spec string, ij IJob) error
}

type Cron struct {
	mu   *sync.Mutex
	cron *cron.Cron
	jobs []*job
}

func NewCron() *Cron {
	return defaultCron
}

var defaultCron = newCron()

func newCron() *Cron {
	return &Cron{
		mu:   &sync.Mutex{},
		cron: cron.New(cron.WithLogger(cron.DiscardLogger)),
	}
}

func (c *Cron) RegisterFunc(spec string, fn func(ctx context.Context) error) error {
	return c.RegisterJob(spec, jobFunc(fn))
}

func (c *Cron) RegisterJob(spec string, ij IJob) error {
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
