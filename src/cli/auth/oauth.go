package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/devhindo/x/src/cli/config"
	"github.com/devhindo/x/src/cli/utils"
)

const (
	RedirectURI = "http://localhost:3000/callback"
	AuthURL     = "https://twitter.com/i/oauth2/authorize"
	TokenURL    = "https://api.twitter.com/2/oauth2/token"
	Scopes      = "tweet.read tweet.write users.read follows.read follows.write offline.access"
)

func StartAuthFlow(app *config.App) error {
	state, err := generateRandomString(32)
	if err != nil {
		return err
	}
	codeVerifier, err := generateCodeVerifier()
	if err != nil {
		return err
	}
	codeChallenge := generateCodeChallenge(codeVerifier)

	// Start local server
	codeChan := make(chan string)
	errChan := make(chan error)
	server := &http.Server{Addr: ":3000"}

	go func() {
		http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			queryState := r.URL.Query().Get("state")
			if queryState != state {
				http.Error(w, "Invalid state", http.StatusBadRequest)
				errChan <- fmt.Errorf("invalid state")
				return
			}
			code := r.URL.Query().Get("code")
			if code == "" {
				http.Error(w, "Code not found", http.StatusBadRequest)
				errChan <- fmt.Errorf("code not found")
				return
			}
			fmt.Fprintf(w, "Authorization successful! You can close this window now.")
			codeChan <- code
		})
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// Construct Auth URL
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", app.ClientID)
	params.Add("redirect_uri", RedirectURI)
	params.Add("scope", Scopes)
	params.Add("state", state)
	params.Add("code_challenge", codeChallenge)
	params.Add("code_challenge_method", "S256")
	authURL := fmt.Sprintf("%s?%s", AuthURL, params.Encode())

	fmt.Println("Opening browser to authenticate...")
	utils.OpenBrowser(authURL)

	// Wait for code or error
	var code string
	select {
	case code = <-codeChan:
	case err := <-errChan:
		return err
	case <-time.After(5 * time.Minute):
		return fmt.Errorf("timeout waiting for authentication")
	}

	// Shutdown server
	server.Shutdown(context.Background())

	// Exchange code for token
	token, err := exchangeCodeForToken(app, code, codeVerifier)
	if err != nil {
		return err
	}

	// Update app with user token
	// We need to reload config to get the latest state (though we passed app pointer)
	// Ideally we should update the config object.
	app.User = token

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Find and update the app in the config
	cfg.AddApp(*app) // AddApp updates if exists

	return config.SaveConfig(cfg)
}

func exchangeCodeForToken(app *config.App, code, codeVerifier string) (*config.User, error) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", app.ClientID)
	data.Set("redirect_uri", RedirectURI)
	data.Set("code_verifier", codeVerifier)

	req, err := http.NewRequest("POST", TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(app.ClientID, app.ClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get token: %s", string(body))
	}

	var result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &config.User{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(result.ExpiresIn) * time.Second),
	}, nil
}

func RefreshToken(app *config.App) error {
	data := url.Values{}
	data.Set("refresh_token", app.User.RefreshToken)
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", app.ClientID)

	req, err := http.NewRequest("POST", TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(app.ClientID, app.ClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to refresh token: %s", string(body))
	}

	var result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	app.User.AccessToken = result.AccessToken
	if result.RefreshToken != "" {
		app.User.RefreshToken = result.RefreshToken
	}
	app.User.Expiry = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	cfg.AddApp(*app)
	return config.SaveConfig(cfg)
}

func generateRandomString(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func generateCodeVerifier() (string, error) {
	const (
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~"
		min   = 43
		max   = 128
	)

	length, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return "", err
	}
	length.Add(length, big.NewInt(min))

	b := make([]byte, length.Int64())
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		b[i] = chars[n.Int64()]
	}

	return string(b), nil
}

func generateCodeChallenge(verifier string) string {
	h := sha256.New()
	h.Write([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
