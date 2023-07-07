package cli

import (
	"github.com/anguloc/zet/cmd/cli/multiple"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "command line interface",
		Long:  "command line interface",
	}

	cmd.AddCommand(multiple.Cmd())

	return cmd
}
