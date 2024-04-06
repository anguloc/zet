package nscan

import (
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nscan",
		Aliases: []string{"nscan"},
		Short:   "有个服务的内网ip、端口找不到了，搞个工具扫描下",
		Long:    "有个服务的内网ip、端口找不到了，搞个工具扫描下",
		Run:     Run,
	}
	return cmd
}
