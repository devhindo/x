package auth

import (
	//"os"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

type data struct {
	Key string `json:"key"`
}

func Verify() bool {
	key, err := lock.ReadLicenseKeyFromFile()
	if err != nil {
		fmt.Println("you are not authenticated | try 'x auth'")
		os.Exit(1)
	}

	k := data{Key: key}

	url := "https://x-blush.vercel.app/api/auth/verify"

	status := postL(url, k)

	if status != 200 {
		fmt.Println("you are not authenticated | try 'x auth'")
		return false
	} else {
		fmt.Println("you're authenticated.")
		return true
	}
}

func postL(url string, k data) int {

    jsonBytes, err := json.Marshal(k)
    if err != nil {
        panic(err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    return resp.StatusCode
}