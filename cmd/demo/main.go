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
)

func main() {
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

func r(ctx context.Context) {
	wg := &sync.WaitGroup{}
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, []os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}...)
	go func() {
		s := <-sig
		go func() {
			_ = app.stopWorker(ctx, s != syscall.SIGQUIT)
		}()
		<-sig
		os.Exit(128 + int(s.(syscall.Signal)))
	}()
	wg.Wait()
}
