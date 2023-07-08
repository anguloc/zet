package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/anguloc/zet/internal/pkg/console"
	"os"

	"github.com/anguloc/zet/internal/app/worker"
	"github.com/anguloc/zet/internal/pkg/application"
	"github.com/anguloc/zet/internal/pkg/conf"
	"github.com/anguloc/zet/internal/pkg/safe"
)

var config string

func main() {
	var err error
	ctx := context.TODO()

	flag.StringVar(&config, "config", "conf/conf.yml", "config path")
	flag.Parse()

	config = safe.Path(config)
	if err = conf.Init(config); err != nil {
		console.Error(err)
		return
	}

	app := application.New()

	// app.RegisterWorker("foo", &worker.FooWorker{})
	// app.RegisterWorker("bar", &worker.BarWorker{})
	app.RegisterWorker("bar", &worker.AnimationRss{})

	if err = app.Init(ctx, ""); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
