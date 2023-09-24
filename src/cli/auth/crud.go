package auth

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
    "encoding/json"
)

func Post(url string, u User) int {

    jsonBytes, err := json.Marshal(u)
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

func GET(url string) {
	// Create a new HTTP request object.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the request and receive the response.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the response body.
	defer resp.Body.Close()

	// Print the response body.
	fmt.Println(string(body))

}
