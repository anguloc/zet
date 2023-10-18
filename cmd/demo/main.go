package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
	"golang.org/x/sync/semaphore"
)

func main() {
	ctx := context.TODO()

	wg := &sync.WaitGroup{}
	t := semaphore.NewWeighted(3)

	list := getList()

	for _, s := range list {
		wg.Add(1)
		if aerr := t.Acquire(ctx, 1); aerr != nil {
			wg.Done()
			continue
		}
		go func(s string) {
			defer wg.Done()
			defer t.Release(1)
			httpRes := mockHttp(s)
			fmt.Printf("%s,res:%s\n", time.Now().Format(time.RFC3339), httpRes)
		}(s)
	}

	wg.Wait()

	return
	console.SetLevel(console.DebugLevel, console.InfoLevel, console.WarnLevel, console.ErrorLevel)
	app := application.New()

	app.RegisterWorker("script", application.NewScript(handle))

	_ = app.Init(ctx)

	if err := app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getList() []string {
	res := make([]string, 0, 10000)
	for i := 0; i < 8000; i++ {
		res = append(res, "zcasd")
	}
	return res
}

func mockHttp(s string) string {
	time.Sleep(1 * time.Second)
	return "asd"
}

func handle(ctx context.Context, param *application.ScriptParam) {
	for param.IsRun() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(param.IsRun())
	}
	time.Sleep(100 * time.Millisecond)
}

func touchPanic() {
	var err error
	res, _ := func() (string, error) {
		defer func() {
			if perr := recover(); perr != nil {
				err = fmt.Errorf("asdasdasd,%+v", perr)
			}
		}()
		panic("asdasda")
	}()
	_ = res
	fmt.Println(err)
}
