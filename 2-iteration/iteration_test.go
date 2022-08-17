package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	var tests = []struct {
		char, want string
	}{
		{"a", "aaaaa"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%q 5 times =%q", test.char, test.want)

		t.Run(testName, func(t *testing.T) {
			got := Repeat(test.char)
			if got != test.want {
				t.Errorf("got %q want %q", got, test.want)
			}
		})
	}
}
