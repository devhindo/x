package auth

import (
	"fmt"
	//"os"
	"github.com/devhindo/x/src/cli/lock"
)

// func check_authentication() {}

func Auth() {

	u := newUser()
	fmt.Println(u.License)
	u.add_user_to_db()
	u.open_browser_to_auth_url()
}

func isAuthenticated() bool {
	key, err := lock.ReadLicenseKeyFromFile()
	if err != nil {
		return false
	}
	if key == "" {
		return false
	}
	return true
}

/*
	if !isAuthenticated() {
		fmt.Println("You are not authenticated.")
		os.Exit(1)
	}
*/