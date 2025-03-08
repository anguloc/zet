package jmcomic

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jmcomic",
		Short: "jmcomic解析",
		Long:  "jmcomic图片解析",
		Run:   Run,
	}

	cmd.Flags().Uint64P("chapter_id", "c", 0, "链接中id")
	cmd.Flags().StringP("img_path", "p", "", "图片地址")

	cmd.Flags().StringP("dir", "d", "", "扫描、输出目录，优先级没有p高")

	return cmd
}
