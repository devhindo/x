package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "x",
	Short: "x is a CLI tool for posting tweets",
	Long: `x is a CLI tool for posting tweets.
	You can post tweets now or in the future.
	You can also clear your credentials and start over.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}	
}

func init() {
	
}