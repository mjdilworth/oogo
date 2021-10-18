package oogo

import "testing"

func TestConfig(t *testing.T) {
	want := "oogo config."
	if got := Config(); got != want {
		t.Errorf("Config() = %q, want %q", got, want)
	}
}
