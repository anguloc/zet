package worker

import (
	"context"
	"fmt"
	"github.com/anguloc/zet/internal/pkg/log"
)

type AnimationRss struct {
	isRunning bool
}

func (w *AnimationRss) Init(ctx context.Context) error {
	fmt.Println("init animation rss")
	return nil
}

func (w *AnimationRss) Run(ctx context.Context) error {
	log.Info(ctx, "AnimationRss_Start")
	ch := make(chan struct{})
	w.isRunning = true
	go func() {
		for w.isRunning {

		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		w.isRunning = false
	}
	<-ch

	log.Info(ctx, "AnimationRss_End")

	return nil
}

func (w *AnimationRss) handle() {
	// 查询rss svr
	// 获取所有关键词
	// 匹配资源
	// 异步获取
}
