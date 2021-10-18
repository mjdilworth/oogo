package auth

import "testing"

func TestAuth(t *testing.T) {
	want := "client Auth"
	if got := Auth(); got != want {
		t.Errorf("Auth() = %q, want %q", got, want)
	}
}
