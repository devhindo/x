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

type License struct {
	License string `json:"license"`
}

type response struct {
    Auth_url string `json:"auth_url"`
}

func postL(url string, l License) {

    jsonBytes, err := json.Marshal(l)
    if err != nil {
        panic(err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

    if err != nil {
        panic(err)
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
		fmt.Println(r.Auth_url)
}