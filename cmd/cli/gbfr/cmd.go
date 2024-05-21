package gbfr

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gbfr",
		Short: "玩玩碧蓝幻想Relink",
		Long:  "玩玩碧蓝幻想Relink",
		Run:   Run,
	}

	cmd.Flags().IntP("display", "", 0, "操作的屏幕序号，从0开始，默认为0")

	return cmd
}
