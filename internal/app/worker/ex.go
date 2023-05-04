package worker

import (
	"context"
	"fmt"
	"time"
)

type FooWorker struct {
}

func (f *FooWorker) Init(ctx context.Context) error {
	fmt.Println("foo init")
	return nil
}

func (f *FooWorker) Run(ctx context.Context) error {
	ch := make(chan struct{})
	isRunning := true
	go func() {
		for isRunning {
			time.Sleep(time.Second)
			fmt.Println("foo running", time.Now().Format("20060102 15:04:05"))
		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		isRunning = false
	}
	<-ch

	fmt.Println("foo start end")
	return nil
}

type BarWorker struct {
}

func (b *BarWorker) Init(ctx context.Context) error {
	fmt.Println("bar init")
	return nil
}

func (b *BarWorker) Run(ctx context.Context) error {
	ch := make(chan struct{})
	isRunning := true
	go func() {
		for isRunning {
			time.Sleep(time.Second)
			fmt.Println("bar running", time.Now().Format("20060102 15:04:05"))
		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		isRunning = false
	}
	<-ch

	fmt.Println("bar start end")
	return nil
}
