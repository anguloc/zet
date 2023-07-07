package multiple

import (
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "multiple -c count -n num [-S shell] [-y yes] command",
		Aliases: []string{"m"},
		Short:   "multiple cli",
		Long:    "multiple cli",
		Run:     Run,
	}
	cmd.Flags().IntP("count", "c", 0, "调用总量")
	cmd.Flags().IntP("num", "n", 0, "goroutine数量")
	cmd.Flags().StringP("shell", "S", "", "命令行执行类型,可选bash、sh")
	cmd.Flags().BoolP("yes", "y", false, "调用总量")

	return cmd
}

