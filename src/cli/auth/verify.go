package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/devhindo/x/src/cli/lock"
)

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

	if status == 200 {
		fmt.Println("you're authenticated.")
		return true
	} else if status == 500 {
		fmt.Println("license not found. try `x auth`")
		return false
	} else if status == 501 {
		fmt.Println("you haven't authorized x cli yet. try `x auth --url`")
		return false
	}
	return true
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