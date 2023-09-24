package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/devhindo/x/src/cli/clear"
	"github.com/devhindo/x/src/cli/lock"
)

func (u *User) generate_auth_url() {
	auth_url := ""
	auth_scopes := "tweet.read%20tweet.write%20users.read%20users.read%20follows.read%20follows.write%20offline.access"
	auth_url += "https://twitter.com/i/oauth2/authorize?response_type=code&client_id="
	auth_url += "emJHZzZHMUdHMF9QRlRIdk45QjY6MTpjaQ"
	redirect_url := "https://x-blush.vercel.app/api/auth"
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

	url := "https://x-blush.vercel.app/api/user/url"

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
