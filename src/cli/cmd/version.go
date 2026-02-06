package cmd

import (
	"github.com/devhindo/x/src/cli/x"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		x.Version()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
