package application

import (
	"context"
	"time"
)

type Script struct {
	Fn func(ctx context.Context, param *ScriptParam)
}

type ScriptParam struct {
	isRun     *bool     // 是否还在运行
	StartTime time.Time // 开始时间
	EndTime   time.Time // 结束时间
}

func (s *ScriptParam) IsRun() bool {
	return *s.isRun == true
}

func NewScript(fn func(ctx context.Context, param *ScriptParam)) *Script {
	return &Script{
		Fn: fn,
	}
}

func (s *Script) Init(_ context.Context) error {
	return nil
}

func (s *Script) Run(ctx context.Context) error {
	if s.Fn == nil {
		return nil
	}
	var isRun = true
	param := &ScriptParam{
		isRun:     &isRun,
		StartTime: time.Now(),
	}
	go func() {
		select {
		case <-ctx.Done():
			isRun = false
		}
	}()
	s.Fn(ctx, param)
	return nil
}
