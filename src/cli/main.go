package main

import (
	//"github.com/devhindo/x/src/cli/x"
	"net/http"
	"fmt"
	"io/ioutil"
)

type User struct {
	Username string
	Secret string
}

func main() {
	//x.Run()
	url := "http://localhost:3000/api/learn"

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






