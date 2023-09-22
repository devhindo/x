package x

import (
	"os"
	"fmt"
	"github.com/devhindo/x/src/cli/help"
	"github.com/devhindo/x/src/cli/auth"
	"github.com/devhindo/x/src/cli/tweet"
)

func HandleArgs() {
	checkArgs()
	switch os.Args[1] {
		case "help":
			help.Help()
		case "auth":
			if len(os.Args) == 2 {
				auth.Auth()
			} else if len(os.Args) == 3 && (os.Args[2] == "--verify" || os.Args[2] == "-v") {
				auth.Verify()
			} else {
				fmt.Println("Unknown command | try 'x help'")
				os.Exit(0)
			}
		case "-t":
			tweet.POST_tweet(os.Args[2])
		default:
			fmt.Println("Unknown command | try 'x help'")
			os.Exit(0)
	}
}

func checkArgs() {
	if len(os.Args) < 2 {
		fmt.Println("No command given | try 'x help'")
		os.Exit(0)
	}
}

