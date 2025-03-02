package cli

import (
	"github.com/anguloc/zet/cmd/cli/click"
	"github.com/anguloc/zet/cmd/cli/file_svr"
	"github.com/anguloc/zet/cmd/cli/gbfr"
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
	cmd.AddCommand(click.Cmd())
	cmd.AddCommand(gbfr.Cmd())
	cmd.AddCommand(file_svr.Cmd())

	return cmd
}
