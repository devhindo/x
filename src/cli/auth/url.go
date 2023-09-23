package auth

import (
	"fmt"
	"os"
	"io"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/devhindo/x/src/cli/env"
	"github.com/devhindo/x/src/cli/lock"
	"github.com/devhindo/x/src/cli/clear"
)

func (u *User) generate_auth_url() {
	env.Load()
	auth_url := ""
	auth_scopes := "tweet.read%20tweet.write%20users.read%20users.read%20follows.read%20follows.write%20offline.access"
	client_id := os.Getenv("CLIENT_ID")
	auth_url += "https://twitter.com/i/oauth2/authorize?response_type=code&client_id="
	auth_url += client_id
	redirect_url := "http://localhost:3000/api/auth"
	auth_url += "&redirect_uri=" + redirect_url
	auth_url += "&scope=" + auth_scopes
	auth_url += "&state=" + u.State + "&code_challenge=" + u.Code_challenge + "&code_challenge_method=S256"
	u.Auth_URL = auth_url
}

func Get_url_db() {
	get_url_from_db()
}




func get_url_from_db() {

	lic, err := lock.ReadLicenseKeyFromFile()
	if err != nil {
		fmt.Println("couldn't generate url | try 'x auth'")
		return
	}

	l := clear.License{
		License: lic,
	}

	url := "http://localhost:3000/api/user/url"

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	req.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var response struct {
        AuthURL string `json:"auth_url"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        panic(err)
    }

    fmt.Println(response.AuthURL)
	
}

