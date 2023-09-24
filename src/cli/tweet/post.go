package tweet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/devhindo/x/src/cli/lock"
)

type Tweet struct {
	License string `json:"license"`
	Tweet   string `json:"tweet"`
}

func POST_tweet(t string) {

	license, err := lock.ReadLicenseKeyFromFile()

	if err != nil {
		fmt.Println("you are not authenticated | try 'x auth'")
		os.Exit(1)
	}

	url := "https://x-blush.vercel.app/api/tweets/post"
	tweet := Tweet{
		License: license,
		Tweet:   t,
	}
	// Create a new HTTP request object.
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBytes, err := json.Marshal(tweet)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Handle the response
	status := resp.StatusCode
	if status == 200 {
		fmt.Println("tweet posted!")
	} else if status == 401 {
		fmt.Println("couldn't get access token | try 'x auth'")
	}
}
