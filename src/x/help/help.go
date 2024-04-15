package help

import (
	"fmt"
)

func Help() {
	fmt.Println()
	fmt.Println("interact with x (twitter) from terminal.")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  x <command>")
	fmt.Println()
	fmt.Println("Commands")
	fmt.Println("  -h       	 show this help")
	fmt.Println("  auth           start authorizing your X account")
	fmt.Println("  auth --url     get auth url if it didn't open browser after running 'x auth'")
	fmt.Println("  auth -v        verify authorization after running 'x auth'")
	fmt.Println("  -t \"text\"   	 post a tweet")
	fmt.Println("  -v             show version")
	fmt.Println("  -c       	 clear authorized account")
	fmt.Println()
	fmt.Println("LEARN MORE")
	fmt.Println("  Cheack source code at: https://github.com/devhindo/x")
}