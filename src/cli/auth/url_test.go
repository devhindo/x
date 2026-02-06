package auth

import (
	"strings"
	"testing"
)

func TestGenerateAuthUrl(t *testing.T) {
	u := &User{
		State:          "test_state",
		Code_challenge: "test_challenge",
	}

	u.generate_auth_url()

	if u.Auth_URL == "" {
		t.Error("generate_auth_url() resulted in empty Auth_URL")
	}

	expectedParts := []string{
		"https://twitter.com/i/oauth2/authorize",
		"response_type=code",
		"client_id=emJHZzZHMUdHMF9QRlRIdk45QjY6MTpjaQ",
		"redirect_uri=https://x-blush.vercel.app/api/auth",
		"scope=tweet.read%20tweet.write%20users.read%20users.read%20follows.read%20follows.write%20offline.access",
		"state=test_state",
		"code_challenge=test_challenge",
		"code_challenge_method=S256",
	}

	for _, part := range expectedParts {
		if !strings.Contains(u.Auth_URL, part) {
			t.Errorf("Auth_URL missing part: %s", part)
		}
	}
}
