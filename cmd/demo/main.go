package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
)

type ADemo struct {
}

func (A ADemo) Init(ctx context.Context) error {
	return nil
}

func (A ADemo) Run(ctx context.Context) error {
	var isStop = false
	go func() {
		select {
		case <-ctx.Done():
			isStop = true
		}
	}()
	for !isStop {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(1)
	}

	return nil
}

func main() {
	console.SetLevel(console.DebugLevel, console.InfoLevel, console.WarnLevel, console.ErrorLevel)
	ctx := context.TODO()
	app := application.New()

	app.RegisterWorker("ADemo", &ADemo{})

	_ = app.Init(ctx)

	if err := app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return

	ctx, cancel := context.WithCancel(context.TODO())
	wg := &sync.WaitGroup{}

	select {
	case <-ctx.Done():

	}

	r(wg, cancel)

	fmt.Println("start")
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err == io.EOF {
				time.Sleep(200 * time.Millisecond)
				continue
			} else {
				break
			}
		}
		if !json.Valid(line) {
			//fmt.Println("no_json")
		}
		//fmt.Println(string(line))
	}

	fmt.Println("end")
}

func r(wg *sync.WaitGroup, cancel func()) {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, []os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}...)
	go func() {
		s := <-sig
		go cancel()
		<-sig
		os.Exit(128 + int(s.(syscall.Signal)))
	}()
	wg.Wait()
}
