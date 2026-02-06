package auth

import (
	"encoding/base64"
	"testing"
)

func TestHashSha256(t *testing.T) {
	input := "hello world"
	// echo -n "hello world" | openssl sha256 -binary | base64
	// uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=
	expectedBase64 := "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="

	got := hash_sha256(input)
	gotBase64 := base64.StdEncoding.EncodeToString(got)

	if gotBase64 != expectedBase64 {
		t.Errorf("hash_sha256(%q) = %q, want %q", input, gotBase64, expectedBase64)
	}
}

func TestGenerateCodeVerifier(t *testing.T) {
	u := &User{}
	u.generate_code_verifier()

	if len(u.Code_verifier) < 43 || len(u.Code_verifier) > 128 {
		t.Errorf("generate_code_verifier() length = %d, want between 43 and 128", len(u.Code_verifier))
	}

	// Check for invalid characters
	allowed := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~"
	for _, char := range u.Code_verifier {
		isAllowed := false
		for _, a := range allowed {
			if char == a {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			t.Errorf("generate_code_verifier() contains invalid character: %c", char)
		}
	}
}

func TestGenerateCodeChallenge(t *testing.T) {
	u := &User{}
	// Set a fixed verifier for reproducibility
	u.Code_verifier = "hello_world_verifier_1234567890" 
	
	// sha256("hello_world_verifier_1234567890")
	// echo -n "hello_world_verifier_1234567890" | openssl sha256 -binary | base64
	// hash = 84c3c33379967676e828114f851080d859e3557451965b1285268c375531d041
	// Base64URL(hash) (without padding)
	
	u.generate_code_challenge()

	// Verify the result
	// The implementation calls hash_sha256 then Base64 RawURL encoding
	hash := hash_sha256(u.Code_verifier)
	expected := base64.RawURLEncoding.EncodeToString(hash)

	if u.Code_challenge != expected {
		t.Errorf("generate_code_challenge() = %q, want %q", u.Code_challenge, expected)
	}
}

func TestGenerateState(t *testing.T) {
	u := &User{}
	length := 127
	u.generate_state(length)

	// The implementation generates random bytes of 'length' then Base64 URL encodes them.
	// So the resulting string length will be roughly length * 4/3.
	
	if u.State == "" {
		t.Error("generate_state() produced empty state")
	}
	
	// Decode back to check byte length
	decoded, err := base64.URLEncoding.DecodeString(u.State)
	if err != nil {
		t.Errorf("generate_state() produced invalid base64: %v", err)
	}
	
	if len(decoded) != length {
		t.Errorf("generate_state() decoded length = %d, want %d", len(decoded), length)
	}
}
