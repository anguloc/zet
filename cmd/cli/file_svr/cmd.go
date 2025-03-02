package file_svr

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "file_server",
		Aliases: []string{"fsvr"},
		Short:   "文件需要在多设备设备传输，现在的工具都需要下载app，smb有时候不太方便",
		Long:    "文件需要在多设备设备传输，现在的工具都需要下载app，smb有时候不太方便",
		Run:     Run,
	}

	cmd.Flags().IntP("port", "p", 41230, "监听的端口")

	return cmd
}
