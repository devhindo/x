package auth

import (
	"fmt"
	"os"

	"github.com/devhindo/x/src/cli/lock"
)

func Get_url_db() {
	l, err := lock.ReadLicenseKeyFromFile()
	if err != nil {
		fmt.Println("you are not authenticated | try 'x auth'")
		os.Exit(1)
	}
	
	url := "https://x-blush.vercel.app/api/user/url"
	k := License{License: l}

	postL(url, k)
}
