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

type data struct {
	Key string `json:"key"`
}

func Verify() bool {
	key, err := lock.ReadLicenseKeyFromFile()
	if err != nil {
		fmt.Println("you are not authenticated | try 'x auth'")
		os.Exit(1)
	}

	k := data{Key: key}

	url := "https://x-blush.vercel.app/api/auth/verify"

	postL(url, k)

	
	return true
}

type response struct {
    Message string `json:"message"`
}

func postL(url string, k data) {

    jsonBytes, err := json.Marshal(k)
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

		//Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)
}