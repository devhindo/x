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
	
	postT(url, tweet)
}

type response struct {
    Message string `json:"message"`
}

func postT(url string, t Tweet) {

    jsonBytes, err := json.Marshal(t)
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
		fmt.Println(r.Message)
}