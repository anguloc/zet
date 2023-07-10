package server

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server [-c config]",
		Aliases: []string{"svr"},
		Short:   "svr",
		Long:    "server",
		Run:     Run,
	}

	cmd.Flags().StringP("config", "c", "conf/conf.yml", "config path")

	return cmd
}
