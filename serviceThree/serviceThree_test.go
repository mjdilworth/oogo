package serviceThree

import "testing"

func TestHello(t *testing.T) {
	want := "serviceThree Hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
