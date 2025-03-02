package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ZSig struct {
	Signal      []os.Signal   // 监听的退出信号量，有默认值
	QuitTimeout time.Duration // 退出等待时间，默认5秒

	// 几个场景打印，可以做成事件，简单搞就不搞那么复杂了，堆在这里吧
	RecvQuitSigPrint func(s os.Signal)
	PanicPrint       func(s any)
	StartPrint       func()
	EndPrint         func()
	ErrEndPrint      func()

	isCancel bool // 是否已经收到取消信号
}

func NewZSig() *ZSig {
	return &ZSig{
		Signal:      []os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM},
		QuitTimeout: time.Second * 5,

		RecvQuitSigPrint: func(s os.Signal) { fmt.Printf("[%s]收到退出信号[%v]\n", time.Now().Format(time.DateTime), s) },
		PanicPrint:       func(s any) { fmt.Printf("[%s]执行异常[%v]\n", time.Now().Format(time.DateTime), s) },
		StartPrint:       func() { fmt.Printf("[%s]开始执行\n", time.Now().Format(time.DateTime)) },
		EndPrint:         func() { fmt.Printf("[%s]结束执行\n", time.Now().Format(time.DateTime)) },
		ErrEndPrint:      func() { fmt.Printf("[%s]超时结束执行\n", time.Now().Format(time.DateTime)) },
	}
}

func (z *ZSig) Wait(ctx context.Context, fn func(ctx context.Context)) {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, z.Signal...)

	ctx, cancel := context.WithCancel(ctx)

	// 信号监听
	go func() {
		s := <-sig
		z.RecvQuitSigPrint(s)
		cancel() // 取消ctx
		z.isCancel = true
		<-sig
		// 二次收到信号，强行退出
		os.Exit(128 + int(s.(syscall.Signal)))
	}()

	// 退出超时以及退出完成监听
	newCtx, completeCancel := context.WithCancel(context.Background())

	z.StartPrint()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				z.PanicPrint(r)
			}
		}()
		fn(ctx)
		cancel()
		completeCancel()
	}()

	<-ctx.Done()

	select {
	case <-newCtx.Done():
		z.EndPrint()
	case <-time.After(z.QuitTimeout):
		z.ErrEndPrint()
	}
}

func (z *ZSig) IsStop() bool {
	return z.isCancel
}
