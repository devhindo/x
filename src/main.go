package main

import (
	//"context"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	//"github.com/michimani/gotwi"
	//"github.com/michimani/gotwi/fields"
	//"github.com/michimani/gotwi/user/userlookup"
	//"github.com/michimani/gotwi/user/userlookup/types"
	//"golang.org/x/tools/godoc/redirect"

	"os"

	"github.com/joho/godotenv"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {

	// Set Twitter API Key
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	client_id := os.Getenv("CLIENT_ID")

	auth_url := ""
	auth_scopes := "tweet.read%20tweet.write%20users.read%20users.read%20follows.read%20follows.write%20offline.access"
	redirect_url := "https://github.com/devhindo/x"

	auth_url += "https://twitter.com/i/oauth2/authorize?response_type=code&client_id="
	auth_url += client_id
	auth_url += "&redirect_uri=" + redirect_url

	auth_url += "&scope=" + auth_scopes

	code_challenge := generateRandomString(12)
	state := generateRandomString(12)
	auth_url += "&state=" + state + "&code_challenge=" + code_challenge + "&code_challenge_method=plain"

	fmt.Println(auth_url)

	//os.Getenv("TWITTER_API_KEY")
	//os.Getenv("TWITTER_API_SECRET_KEY")

	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}


}

func run(ctx context.Context) error {
	authToken := os.Getenv("NGROK_AUTHTOKEN")
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtoken(authToken),
	)
	if err != nil {
		return err
	}

	log.Println("tunnel created:", tun.URL())

	return http.Serve(tun, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ngrok-go!")
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
