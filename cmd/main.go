package main

import (
	"fmt"
	"github.com/anguloc/zet/cmd/handle"
	"github.com/anguloc/zet/cmd/rss"
	"github.com/spf13/cobra"
	"os"
)

var root = cobra.Command{
	Use:     "zet",
	Version: "0.0.1",
}

func main() {

	root.AddCommand(handle.Cmd)
	root.AddCommand(rss.Cmd)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
