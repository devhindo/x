package x

import (
	"os"
	"fmt"
	"github.com/devhindo/x/src/cli/help"
	"github.com/devhindo/x/src/cli/auth"
)

func HandleArgs() {
	checkArgs()
	switch os.Args[1] {
		case "help":
			help.Help()
		case "auth":
			//Auth()
		case "tweet":
			//Tweet()
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
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments | try 'x help'")
		os.Exit(0)
	}
}

