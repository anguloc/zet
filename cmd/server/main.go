package main

import (
	"context"
	"fmt"
	"os"

	"github.com/anguloc/zet/internal/app/worker"
	"github.com/anguloc/zet/internal/pkg/application"
)

func main() {
	var err error
	ctx := context.TODO()

	app := application.New()

	app.RegisterWorker("foo", &worker.FooWorker{})
	app.RegisterWorker("bar", &worker.BarWorker{})

	if err = app.Init(ctx, ""); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
