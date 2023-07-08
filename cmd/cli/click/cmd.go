package click

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "click",
		Aliases: []string{"c"},
		Short:   "click",
		Long:    "自动点击",
		Run:     Run,
	}

	return cmd
}
