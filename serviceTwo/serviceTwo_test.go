package servicetwo

import "testing"

func TestHello(t *testing.T) {
	want := "serviceTwo Hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
