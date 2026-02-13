package tweet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/devhindo/x/src/cli/auth"
	"github.com/devhindo/x/src/cli/config"
)

type Tweet struct {
	Text string `json:"text"`
}

func POST_tweet(t string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	app := cfg.GetActiveApp()
	if app == nil || app.User == nil {
		fmt.Println("No active app or user not authenticated. Run 'x init use' or 'x auth'.")
		os.Exit(1)
	}

	if time.Now().After(app.User.Expiry) {
		fmt.Println("Token expired, refreshing...")
		if err := auth.RefreshToken(app); err != nil {
			fmt.Println("Error refreshing token:", err)
			os.Exit(1)
		}
	}

	url := "https://api.twitter.com/2/tweets"
	tweet := Tweet{
		Text: t,
	}

	jsonBytes, err := json.Marshal(tweet)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", "Bearer "+app.User.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error posting tweet:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 201 {
		fmt.Printf("Error posting tweet: %s\n", string(body))
		os.Exit(1)
	}

	fmt.Println("Tweet posted successfully! 🐦")
}
