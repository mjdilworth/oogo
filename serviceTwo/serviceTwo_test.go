package clientlib

import "testing"

func TestHello(t *testing.T) {
	want := "client Hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
