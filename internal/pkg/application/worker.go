package application

import (
	"context"
	"fmt"
	"time"
)

type IWorker interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	GracefulStop(ctx context.Context) error
}

type FooWorker struct {
}

func (f *FooWorker) Init(ctx context.Context) error {
	fmt.Println("foo init")
	return nil
}

func (f *FooWorker) Start(ctx context.Context) error {
	for {
		time.Sleep(time.Second)
		fmt.Println("foo running", time.Now().Format("20060102 15:04:05"))
	}

	fmt.Println("foo start end")
	return nil
}

func (f *FooWorker) Stop(ctx context.Context) error {
	fmt.Println("receive foo stop")
	return nil
}

func (f *FooWorker) GracefulStop(ctx context.Context) error {
	fmt.Println("receive foo g stop")
	return nil
}

type BarWorker struct {
}

func (b *BarWorker) Init(ctx context.Context) error {
	fmt.Println("bar init")
	return nil
}

func (b *BarWorker) Start(ctx context.Context) error {
	for {
		time.Sleep(time.Second)
		fmt.Println("bar running", time.Now().Format("20060102 15:04:05"))
	}

	fmt.Println("bar start end")
	return nil
}

func (b *BarWorker) Stop(ctx context.Context) error {
	fmt.Println("receive bar stop")
	return nil
}

func (b *BarWorker) GracefulStop(ctx context.Context) error {
	fmt.Println("receive bar g stop")
	return nil
}
