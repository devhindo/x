package auth

import (
	"crypto/rand"
	"encoding/base64"
	"os"
)

func Auth_url() string {
	return construct_auth_url("https://x-blush.vercel.app/api")

}

func construct_auth_url(redirect_url string) string {
	auth_url := ""
	auth_scopes := "tweet.read%20tweet.write%20users.read%20users.read%20follows.read%20follows.write%20offline.access"
	client_id := os.Getenv("CLIENT_ID")
	auth_url += "https://twitter.com/i/oauth2/authorize?response_type=code&client_id="
	auth_url += client_id
	auth_url += "&redirect_uri=" + redirect_url
	auth_url += "&scope=" + auth_scopes
	code_challenge := generateRandomString(12)
	state := generateRandomString(12)
	auth_url += "&state=" + state + "&code_challenge=" + code_challenge + "&code_challenge_method=plain"
	return auth_url
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// return only characters in base64 alphabet, removes characters that are not url safe
	return base64.RawURLEncoding.EncodeToString(b)
}