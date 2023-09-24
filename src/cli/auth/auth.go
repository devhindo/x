package auth

import (
	"fmt"
	"os"

	"github.com/devhindo/x/src/cli/lock"
)

// func check_authentication() {}

func Auth() {

	checkIfUserExists()

	u := newUser()
	u.add_user_to_db()
	u.open_browser_to_auth_url()
	fmt.Println("please authorize X CLI in your browser then run 'x auth --verify'")
	fmt.Println("if the browser does not open, run 'x auth --url` to get the authorization url")
}

func checkIfUserExists() {
	_, err := lock.ReadLicenseKeyFromFile()
	if err == nil {
		fmt.Println("a user is already logged in | try 'x -h'")
		os.Exit(0)
	}
}
