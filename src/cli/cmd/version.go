package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("x CLI v1.0.0 (Local)")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
