package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	var tests = []struct {
		char  string
		times int
		want  string
	}{
		{"a", 5, "aaaaa"},
		{"a", 0, ""},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%q %q times=%q", test.char, test.times, test.want)

		t.Run(testName, func(t *testing.T) {
			got := Repeat(test.char, test.times)
			if got != test.want {
				t.Errorf("got %q want %q", got, test.want)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}
