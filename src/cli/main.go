package main

import (
	"os"

	"github.com/devhindo/x/src/cli/cmd"
)

func main() {
	// Support legacy commands starting with dash by rewriting them to their aliased subcommands
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-t":
			os.Args[1] = "t"
		case "-f":
			os.Args[1] = "f"
		case "-v":
			os.Args[1] = "v"
		}
	}

	cmd.Execute()
}

