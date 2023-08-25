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


	

	//os.Getenv("TWITTER_API_KEY")
	//os.Getenv("TWITTER_API_SECRET_KEY")
	ngrok_error, tun_URL := run(context.Background())
	if ngrok_error != nil {
		log.Fatal(err)
	}
	log.Println(tun_URL)
	log.Println(tun_URL)
	log.Println(tun_URL)
	log.Println(tun_URL)
	log.Println(tun_URL)
	// get the ngrok url

	


	

}

func run(ctx context.Context) (error, string) {
	authToken := os.Getenv("NGROK_AUTHTOKEN")
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtoken(authToken),
		// todo set ngrok-skip-browser-warning to anyvalue
	)
	if err != nil {
		return err, ""
	}
	
	log.Println("tunnel created:", tun.URL())

	construct_auth_url(tun.URL())

	return http.Serve(tun, http.HandlerFunc(handler)), tun.URL()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ngrok-go!")
}




func construct_auth_url(redirect_url string) {
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

	fmt.Println(auth_url)
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
