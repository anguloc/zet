package rss

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "rss",
	Short: "rss",
	Long:  "Really Simple Syndication",
	Run:   Run,
}

func init() {
}

func Run(cmd *cobra.Command, args []string) {

}
