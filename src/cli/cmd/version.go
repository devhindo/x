package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

//TODO: make it for -v and --version too

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version of the CLI",
	Long:    `Print the version of the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}