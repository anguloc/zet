package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
)

type Foo struct {
	A struct {
		B struct {
			C string
		}
	}
}

func (f *Foo) Bar() string {
	res := f.A.B.C
	return res
}

func (f *Foo) Qux() string {
	return f.A.B.C
}

func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("Total allocated in bytes:", m.TotalAlloc)
	fmt.Println("Currently allocated in bytes:", m.Alloc)
	fmt.Println("Number of mallocs:", m.Mallocs)
	fmt.Println("Number of frees:", m.Frees)
}

func main() {
	ctx := context.TODO()

	out1 := &bytes.Buffer{}
	out2 := &bytes.Buffer{}

	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/c", "start", "http://www.baidu.com")
	} else if runtime.GOOS == "linux" {
		c = exec.Command("cmd", "/c", "start", "http://www.baidu.com")
	} else {
		fmt.Println("no os")
		return
	}

	c.Stdout = out1
	c.Stderr = out2
	err := c.Start()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(out1)
	fmt.Println(out2)

	return

	f := &Foo{}
	_ = f

	for i := 0; i < 4096; i++ {
		f.A.B.C += "a"
	}

	fmt.Println("-------------------------")
	f.Bar()
	fmt.Println("-------------------------")
	f.Qux()

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

func limit1() {
	//wg := &sync.WaitGroup{}
	//t := semaphore.NewWeighted(3)
	//
	//list := getList()
	//
	//for _, s := range list {
	//	wg.Add(1)
	//	if aerr := t.Acquire(ctx, 1); aerr != nil {
	//		wg.Done()
	//		continue
	//	}
	//	go func(s string) {
	//		defer wg.Done()
	//		defer t.Release(1)
	//		httpRes := mockHttp(s)
	//		fmt.Printf("%s,res:%s\n", time.Now().Format(time.RFC3339), httpRes)
	//	}(s)
	//}
	//
	//wg.Wait()
}
