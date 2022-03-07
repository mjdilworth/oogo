package main

import "testing"

func TestRomanTodecimal(t *testing.T) {

	//table driven test
	var tests = []struct {
		a string
		b int
	}{
		{"i", 1},
		{"ii", 2},
		{"iii", 3},
		{"iv", 4},
		{"v", 5},
		{"vi", 6},
		{"vii", 7},
		{"viii", 8},
		{"ix", 9},
		{"x", 10},
		{"xiii", 13},
		{"xiv", 14},
		{"l", 50},
		{"lxxx", 80},
		{"mcmlxxxviii", 1988},
		{"mcmxcix",1999},

	}

	for _, test := range tests {
		want := test.b
		if got := romanTodecimal(test.a); got != want {
			t.Errorf("romanTodecimal() = %d, want %d", got, want)
		}

	}
}
