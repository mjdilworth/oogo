package serviceone

import "testing"

func TestHello(t *testing.T) {
	want := "serviceOne Hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
