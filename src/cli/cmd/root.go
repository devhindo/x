package cmd

import (
	"os"

	"github.com/devhindo/x/src/cli/tweet"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "x",
	Short: "x CLI - Post tweets from your terminal",
	Long:  `x is a CLI tool that allows you to post tweets to X (formerly Twitter) directly from your terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			// Join args just in case, but usually it's one arg "msg"
			// The original code uses os.Args[1] or args[0]
			tweet.POST_tweet(args[0])
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
