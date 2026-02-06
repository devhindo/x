package cmd

import (
	"github.com/devhindo/x/src/cli/tweet"
	"github.com/spf13/cobra"
)

var tweetCmd = &cobra.Command{
	Use:     "tweet [message]",
	Short:   "Post a tweet",
	Aliases: []string{"t"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tweet.POST_tweet(args[0])
	},
}

func init() {
	rootCmd.AddCommand(tweetCmd)
}
