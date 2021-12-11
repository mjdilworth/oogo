package main

import "testing"

func TestConfig(t *testing.T) {
	want := "running configurations"
	if got := Config(); got != want {
		t.Errorf("Config() = %q, want %q", got, want)
	}
}
