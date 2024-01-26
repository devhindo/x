package tweet

// x t "hi" 5h6m7s

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/devhindo/x/src/cli/lock"
)

type FutureTweet struct {
	License string `json:"license"`
	Tweet   string `json:"tweet"`
	Hours int `json:"hours"`
	Minutes int `json:"minutes"`
}

func PostFutureTweet(c []string) {
	
	
	// x t "hi" 5h6m7s

	tweetText, tweetTime, err := handleFutureTweetArgs(c)
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}

	hrs, mins, err := handleTweetTime(tweetTime)

	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
	
	license, err := lock.ReadLicenseKeyFromFile()
	
	if err != nil {
		fmt.Println("you are not authenticated | try 'x auth'")
		return
	}

	tweet := FutureTweet{
		License: license,
		Tweet: tweetText,
		Hours: hrs,
		Minutes: mins,
	}
	
	url := ""

	err = postFutureTweetToServer(url, tweet)

	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}	
}


func postFutureTweetToServer(url string, t FutureTweet) error {
	return nil
}

func handleFutureTweetArgs(c []string) (string, string, error) {
	if len(c) < 3 {
		return "", "", fmt.Errorf("no tweet given | try 'x f --help'")
	}

	if c[2] == "-h" || c[2] == "--help" {
		fmt.Println("post future tweets")
		fmt.Println("using delayed times in this form")
		fmt.Println("x f \"hi\" 2h3m")
		fmt.Println("h -> hours")
		fmt.Println("m -> minutes")
		fmt.Println("this tweet would be scheduled to be posted after 2 hours and 3 minuets")
		return c[2], c[3], nil
	}

	if len(c) < 4 {
		fmt.Println("No schedule time is given | try 'x f --help'")
	}
	return c[2], c[3], nil
}

func handleTweetTime(t string) (int, int, error) {
	hrs := 0
	mins := 0

	// Check if the string is empty
	if len(t) == 0 {
		return hrs, mins, fmt.Errorf("empty time string")
	}

	containsH := strings.Contains(t, "h")
	containsM := strings.Contains(t, "m")

	if containsH && containsM {
		// Split the string into hours and minutes
		timeParts := strings.Split(t, "h")
		if len(timeParts) != 2 {
			return hrs, mins, fmt.Errorf("invalid time string")
		}

		hours, err := strconv.Atoi(timeParts[0])
		if err != nil {
			return hrs, mins, fmt.Errorf("invalid time string")
		}

		minutes, err := strconv.Atoi(strings.TrimSuffix(timeParts[1], "m"))
		if err != nil {
			return hrs, mins, fmt.Errorf("invalid time string")
		}

		hrs = hours
		mins = minutes
	} else if containsH {
		// Extract the hours from the string
		hours, err := strconv.Atoi(strings.TrimSuffix(t, "h"))
		if err != nil {
			return hrs, mins, fmt.Errorf("invalid time string")
		}

		hrs = hours
	} else if containsM {
		// Extract the minutes from the string
		minutes, err := strconv.Atoi(strings.TrimSuffix(t, "m"))
		if err != nil {
			return hrs, mins, fmt.Errorf("invalid time string")
		}

		mins = minutes
	} else {
		return hrs, mins, fmt.Errorf("invalid time string")
	}

	return hrs, mins, nil
}