package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var tests = []struct {
		name, language, want string
	}{
		{"Jess", "english", "Hello Jess"},
		{"bob", "spanish", "Hola bob"},
		{"Sue", "french", "Bounjour Sue"},
		{"", "", "Hello World"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf(
			"%q%q%q", test.name, test.language, test.want,
		)
		t.Run(testName, func(t *testing.T) {
			got := Hello(test.name, test.language)
			assertCorrectString(t, got, test.want)
		})
	}
}

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
