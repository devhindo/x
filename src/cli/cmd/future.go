package cmd

import (
	"github.com/devhindo/x/src/cli/tweet"
	"github.com/spf13/cobra"
)

var futureCmd = &cobra.Command{
	Use:     "future [message] [time]",
	Short:   "Post a future tweet",
	Aliases: []string{"f"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Construct fake os.Args for compatibility
		fakeArgs := []string{"x", "f", args[0], args[1]}
		tweet.PostFutureTweet(fakeArgs)
	},
}

func init() {
	rootCmd.AddCommand(futureCmd)
}
