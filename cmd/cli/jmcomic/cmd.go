package jmcomic

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jmcomic",
		Short: "jmcomic解析",
		Long:  "jmcomic解析",
		Run:   Run,
	}
	return cmd
}
