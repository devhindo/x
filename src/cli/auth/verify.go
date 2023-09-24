package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io"

	"github.com/devhindo/x/src/cli/lock"
)

type License struct {
	License string `json:"license"`
}

func Verify() bool {
	l, err := lock.ReadLicenseKeyFromFile()
	if err != nil {
		fmt.Println("you are not authenticated | try 'x auth'")
		os.Exit(1)
	}

	k := License{License: l}

	url := "https://x-blush.vercel.app/api/auth/verify"

	postL(url, k)

	
	return true
}

type response struct {
    Message string `json:"message"`
}

func postL(url string, l License) {

    jsonBytes, err := json.Marshal(l)
    if err != nil {
        panic(err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

    if err != nil {
        fmt.Println("can't reach server to verify user")
		os.Exit(0)
    }

    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			panic(err)
		}

		var r response

		err = json.Unmarshal(body, &r)

		//Convert bytes to String and print
		fmt.Println(r.Message)
}