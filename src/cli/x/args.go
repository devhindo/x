package x

import (
	"os"
	"fmt"

	"github.com/devhindo/x/src/cli/help"
	"github.com/devhindo/x/src/cli/auth"
	"github.com/devhindo/x/src/cli/tweet"
	"github.com/devhindo/x/src/cli/clear"

)

func HandleArgs() {
	checkArgs()
	switch os.Args[1] {
		case "help":
			checkArgsequals2()
			help.Help()
		case "--help":
			checkArgsequals2()
			help.Help()
		case "-h":
			checkArgsequals2()
			help.Help()
		case "auth":
			if len(os.Args) == 2 {
				auth.Auth()
			} else if len(os.Args) == 3 && (os.Args[2] == "--verify" || os.Args[2] == "-v") {
				auth.Verify()
			} else if len(os.Args) == 3 && (os.Args[2] == "--clear" || os.Args[2] == "-c") {
				clear.StartOver()
			} else if len(os.Args) == 3 && os.Args[2] == "--url" {
				auth.Get_url_db()
			} else {
				fmt.Println("Unknown command | try 'x help'")
				os.Exit(0)
			}
		case "-t":
			checkArgsequals3()
			fmt.Println(len(os.Args))
			tweet.POST_tweet(os.Args[2])
		case "--tweet":
			checkArgsequals3()
			tweet.POST_tweet(os.Args[2])
		case "--version":
			checkArgsequals2()
			Version()
		case "-v":
			checkArgsequals2()
			Version()
		case "version":
			checkArgsequals2()
			Version()
		default:

			if len(os.Args) != 2 {
				fmt.Println("Unknown command | try 'x help'")
				os.Exit(0)
			}
			
			tweet.POST_tweet(os.Args[1])
			
	}
}

func checkArgs() {
	if len(os.Args) < 2 {
		fmt.Println("No command given | try 'x help'")
		os.Exit(0)
	}
}

func checkArgsequals2() {
	if len(os.Args) != 2 {
		fmt.Println("Unknown command | try 'x help'")
		os.Exit(0)
	}
}

func checkArgsequals3() {
	if len(os.Args) != 3 {
		fmt.Println("Unknown command | try 'x help'")
		os.Exit(0)
	}
}