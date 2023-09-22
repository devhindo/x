package auth

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
    "log"
    "encoding/json"

	"github.com/devhindo/x/src/cli/lock"

)

func POST(url string, u User) {
	// Create a new HTTP request object.
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBytes, err := json.Marshal(u)
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
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Handle the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	status := resp.StatusCode
	
	if status != 200 {
		fmt.Println("error adding user")
	} else {
		err := lock.WriteLicenseKeyToFile(u.License)
		if err != nil {
			fmt.Println("coudln't write license key to file")
			return
		}
	}

    //fmt.Println(string(body))
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
