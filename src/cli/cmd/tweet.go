package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

/*
 x tweet -c "hi" -w 2d5h6m7s -d 2020-12-12 -t 3:34pm
*/

func init() {
	rootCmd.AddCommand(tweetCmd)

	tweetCmd.Flags().StringP("content", "c", "", "Tweet message")
	tweetCmd.Flags().StringP("wait", "w", "0", "When to post the tweet")
	tweetCmd.Flags().StringP("date", "d", "0", "Date to post the tweet")
	tweetCmd.Flags().StringP("time", "t", "0", "Time to post the tweet")
	// tweetCmd.Flags().BoolP("media", "m", false, "Add media to the tweet") TODO: feat: post media content
}

var (
	tweetCmd = &cobra.Command{
		Use:   "tweet",
		Short: "Post a tweet",
		Long:  `Post a tweet.`,
		RunE:  tweetCmdRun,
	}
)

func tweetCmdRun(cmd *cobra.Command, args []string) error {
	content, err := cmd.Flags().GetString("content")
	if err != nil {
		err = fmt.Errorf("error getting content  flag: %v", err)
		return err
	}

	wait, err := cmd.Flags().GetString("wait")
	h, m, s, ms, err := handleWaitArg(wait)
	if err != nil {
		err = fmt.Errorf("error getting wait flag: %v", err)
		return err
	}
	fmt.Println(h, m, s, ms)

	fmt.Println(content)

	return nil
}

func calcWaitingTime() {
	UTCtime := time.Now().UTC()
	fmt.Println(UTCtime)
}

// sra7a ai generated, I wouldn't be able to do that by myself
func handleWaitArg(wait string) (int, int, int, int, error) {
	// Define the regular expression patterns for each unit
	patterns := map[string]*regexp.Regexp{
		"d": regexp.MustCompile(`(\d*)d`),
		"h": regexp.MustCompile(`(\d*)h`),
		"m": regexp.MustCompile(`(\d*)m`),
		"s": regexp.MustCompile(`(\d*)s`),
	}

	// Initialize the time units
	var days, hours, minutes, seconds int

	// Extract each unit from the argument
	for unit, pattern := range patterns {
		match := pattern.FindStringSubmatch(wait)
		if match != nil {
			value, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, 0, 0, 0, fmt.Errorf("invalid waiting time format: %s", wait)
			}

			switch unit {
			case "d":
				days = value
			case "h":
				hours = value
			case "m":
				minutes = value
			case "s":
				seconds = value
			}
		}
	}

	return days, hours, minutes, seconds, nil
}


