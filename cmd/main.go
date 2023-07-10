package main

import (
	"fmt"
	"os"

	"github.com/anguloc/zet/cmd/cli"
	"github.com/anguloc/zet/cmd/handle"
	"github.com/anguloc/zet/cmd/rss"
	"github.com/anguloc/zet/cmd/server"
	"github.com/anguloc/zet/internal/pkg/console"
	"github.com/spf13/cobra"
)

var root = cobra.Command{
	Use:     "zet",
	Version: "0.0.1",
}

func main() {
	console.SetLevel(console.DebugLevel, console.InfoLevel, console.WarnLevel, console.ErrorLevel)

	root.AddCommand(handle.Cmd())
	root.AddCommand(rss.Cmd())
	root.AddCommand(cli.Cmd())
	root.AddCommand(server.Cmd())
	root.SetFlagErrorFunc(func(command *cobra.Command, err error) error {
		console.Debugf("参数错误:%s\n", err)
		return command.Usage()
	})

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
