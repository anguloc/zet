package server

import (
	"fmt"
	"os"

	"github.com/anguloc/zet/internal/app/worker"
	"github.com/anguloc/zet/internal/pkg/application"
	"github.com/anguloc/zet/internal/pkg/conf"
	"github.com/anguloc/zet/internal/pkg/console"
	"github.com/anguloc/zet/internal/pkg/safe"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, _ []string) {
	var err error
	ctx := cmd.Context()

	config, err := cmd.Flags().GetString("")
	if err != nil {
		console.Error("配置文件异常", err)
		return
	}

	config = safe.Path(config)
	if err = conf.Init(config); err != nil {
		console.Error(err)
		return
	}

	app := application.New()

	// app.RegisterWorker("foo", &worker.FooWorker{})
	// app.RegisterWorker("bar", &worker.BarWorker{})
	app.RegisterWorker("dmhy_rss", &worker.DmhyRss{})

	if err = app.Init(ctx, ""); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}