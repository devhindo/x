package auth

import (
	//"os"
	"github.com/devhindo/x/src/cli/lock"
	"bytes"
	"fmt"
	"io"
	"net/http"
    "encoding/json"
	"log"
	"os"
)

// func check_authentication() {}

func Auth() {
	u := newUser()
	u.add_user_to_db()
	u.open_browser_to_auth_url()
}

type data struct{
	Key string `json:"key"`
}

func Verify() bool {
	key, err := lock.ReadLicenseKeyFromFile()
	if err != nil {
		fmt.Println("you are not authenticated | try 'x auth'")
		os.Exit(1)
	}

	k := data{Key: key}

	url := "http://localhost:3000/api/auth/verify"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}


	jsonBytes, err := json.Marshal(k)
	if err != nil {
		panic(err)
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
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}


	
	status := resp.StatusCode

	if status != 200 {
		fmt.Println("you are not authenticated | try 'x auth'")
		return false
	} else {
		fmt.Println("you're authenticated.")
		return true
	}
}

/*
	if !isAuthenticated() {
		fmt.Println("You are not authenticated.")
		os.Exit(1)
	}
*/