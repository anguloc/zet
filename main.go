package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "zet",
	Short:   "private tools",
	Version: "1.0.0",
}

func main() {
	fmt.Println("vim-go")
}
