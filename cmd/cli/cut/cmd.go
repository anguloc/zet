package cut

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cut",
		Short: "截图",
		Long:  "截图",
		Run:   Run,
	}

	cmd.Flags().IntVarP(&x, "x", "x", 0, "横坐标原点")
	cmd.Flags().IntVarP(&y, "y", "y", 0, "纵坐标原点")
	cmd.Flags().IntVarP(&w, "w", "w", 0, "截图宽度")
	cmd.Flags().IntVarP(&h, "h", "h", 0, "截图高度")
	cmd.Flags().StringVarP(&file, "file", "f", "", "保存文件名")

	return cmd
}
