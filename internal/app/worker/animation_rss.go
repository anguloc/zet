package worker

import (
	"context"
	"fmt"
)

type AnimationRss struct {
}

func (w *AnimationRss) Init(ctx context.Context) error {
	fmt.Println("init animation rss")
	return nil
}
func (w *AnimationRss) Run(ctx context.Context) error {
	return nil
}
