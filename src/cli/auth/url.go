package auth

import (
	"os"
	"github.com/devhindo/x/src/cli/env"
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






