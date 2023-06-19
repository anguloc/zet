package rss

import (
	"github.com/anguloc/zet/internal/app/rss"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "rss",
	Short: "rss",
	Long:  "Really Simple Syndication",
	Run:   Run,
}

func init() {
}

func Run(cmd *cobra.Command, args []string) {
	biz := rss.NewDmhy()
	biz.Run(ctx)
}
