package servicelib

import "testing"

func TestHello(t *testing.T) {
	want := "server Hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
