package auth

import "testing"

func TestIsAuthenticated(t *testing.T) {
	got := IsAuthenticated()
	if got != false {
		t.Errorf("IsAuthenticated() = %v, want false", got)
	}
}
