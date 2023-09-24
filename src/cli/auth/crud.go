package auth

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
    "encoding/json"

	"github.com/devhindo/x/src/cli/lock"

)

func POST(url string, u User) {
		
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	status := resp.StatusCode
	fmt.Println(status)

	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
    //		panic(err)
	//}

	//Convert bytes to String and print
	//jsonStr := string(body)
	//fmt.Println("Response: ", jsonStr)

	
	if status != 200 {
		fmt.Println("error adding user")
	} else {
		
		err := lock.WriteLicenseKeyToFile(u.License)
		if err != nil {
			fmt.Println("coudln't write license key to file")
			return
		}
	}

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
