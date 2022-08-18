package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		slice []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 3}, 4},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%d = %d", test.slice, test.want)

		t.Run(testName, func(t *testing.T) {
			got := Sum(test.slice)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
