package tweet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

type Tweet struct {
	State string `json:"state"`
	Tweet string `json:"tweet"`
}

func POST_tweet() {
	url := "http://localhost:3000/api/tweets/post"
	tweet := Tweet{
		State: "#",
		Tweet: "gaaaa",
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(string(body))
}
