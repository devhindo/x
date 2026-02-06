package lock

import (
	"encoding/base64"
	"testing"
)

func TestGenerateLicenseKey(t *testing.T) {
	key, err := GenerateLicenseKey()
	if err != nil {
		t.Fatalf("GenerateLicenseKey() error = %v", err)
	}
	if key == "" {
		t.Error("GenerateLicenseKey() returned empty string")
	}

	// Verify it is valid base64
	_, err = base64.StdEncoding.DecodeString(key)
	if err != nil {
		t.Errorf("GenerateLicenseKey() returned invalid base64: %v", err)
	}
}
