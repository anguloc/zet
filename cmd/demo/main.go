package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
)

func main() {
	console.SetLevel(console.DebugLevel, console.InfoLevel, console.WarnLevel, console.ErrorLevel)
	ctx := context.TODO()
	app := application.New()

	app.RegisterWorker("script", application.NewScript(handle))

	_ = app.Init(ctx)

	if err := app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handle(ctx context.Context, param *application.ScriptParam) {
	for param.IsRun() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(param.IsRun())
	}
	time.Sleep(100 * time.Millisecond)
}
