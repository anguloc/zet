package server

import (
	"fmt"
	"os"

	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/console"
	"github.com/anguloc/zet/pkg/cron"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, _ []string) {
	var err error
	ctx := cmd.Context()

	config, err := cmd.Flags().GetString("config")
	if err != nil {
		console.Error("配置文件异常", err)
		return
	}

	app := application.New()

	app.RegisterWorker("cron", cron.NewWorker())
	// app.RegisterWorker("foo", &worker.FooWorker{})
	// app.RegisterWorker("bar", &worker.BarWorker{})

	if err = app.Init(ctx, config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = app.Run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
