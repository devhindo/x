package main

import (
	//"github.com/devhindo/x/src/cli/x"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
    "log"
    "encoding/json"

)

type User struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

func main() {
	//x.Run()
	url := "http://localhost:3000/api/learn"

	GET(url)

	POST(url, User{
		Username: "devhindo",
		Secret:   "123456",
	})

}

func POST(url string, user User) {
	// Create a new HTTP request object.
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Handle the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(string(body))
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the response body.
	defer resp.Body.Close()

	// Print the response body.
	fmt.Println(string(body))

}
