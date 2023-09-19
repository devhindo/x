package auth

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"github.com/devhindo/x/src/cli/env"
)



func Auth_url() string {
	//return construct_auth_url("https://x-blush.vercel.app/api")
	return construct_auth_url("http://localhost:3000/api/auth")
}

func construct_auth_url(redirect_url string) string {
	env.Load()
	auth_url := ""
	auth_scopes := "tweet.read%20tweet.write%20users.read%20users.read%20follows.read%20follows.write%20offline.access"
	client_id := os.Getenv("CLIENT_ID")
	auth_url += "https://twitter.com/i/oauth2/authorize?response_type=code&client_id="
	auth_url += client_id
	auth_url += "&redirect_uri=" + redirect_url
	auth_url += "&scope=" + auth_scopes
	code_challenge := GenerateRandomString(127)
	state := GenerateRandomString(127)
	auth_url += "&state=" + state + "&code_challenge=" + code_challenge + "&code_challenge_method=plain"
	return auth_url
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// return only characters in base64 alphabet, removes characters that are not url safe
	return base64.RawURLEncoding.EncodeToString(b)
}

func generateURL() string {
	code_verifier := GenerateRandomString(127)
}


