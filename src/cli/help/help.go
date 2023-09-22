package help

import (
	"fmt"
)

func Help() {
	fmt.Println()
	fmt.Println("tweet from terminal and other stuff.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  x <command> <arguments>")
	fmt.Println()
	fmt.Println("Commands")
	fmt.Println("  help          show this help")
	fmt.Println("  auth          authorize your account twitter")
	fmt.Println("  auth -v          verify authorization")
	fmt.Println("  -t <text>  	 post a tweet")
	fmt.Println("  --tweet <text>  	 post a tweet")
	fmt.Println()
	fmt.Println("LEARN MORE")
	fmt.Println("  Cheack source code at: https://github.com/devhindo/x")
}