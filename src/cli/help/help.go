package help

import (
	"fmt"
)

func Help() {
	fmt.Println("tweet from terminal and other stuff.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  x <command> <arguments>")
	fmt.Println()
	fmt.Println("Commands")
	fmt.Println("  help          show this help")
	fmt.Println("  auth         authorize your account twitter")
	fmt.Println("  tweet <text>  tweet text")
	fmt.Println()
	fmt.Println("LEARN MORE")
	fmt.Println("  Cheack source code at: https://github.com/devhindo/x")
}