package cmd

import (
	"github.com/devhindo/x/src/cli/tweet"

	"github.com/spf13/cobra"
)

var (
	wait string
	date string
	content string
)

func init() {
	rootCmd.AddCommand(futureCmd)
}

var futureCmd = &cobra.Command{
	Use:   "-f",
	Short: "Post future tweets",
	Long:  `Post future tweets.`,
	Run:   futureCmdRun,
}

func futureCmdRun(cmd *cobra.Command, args []string) {
	tweet.PostFutureTweet(args)
}
