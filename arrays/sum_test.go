package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		slice []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 3}, 4},
	}

	for _, test := range tests {
		test := test
		testName := fmt.Sprintf("%d = %d", test.slice, test.want)

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got := Sum(test.slice)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		slice1, slice2, want []int
	}{
		{[]int{1, 2}, []int{4, 5}, []int{3, 9}},
		{slice1: []int{1, 3}, want: []int{4, 0}},
	}

	for _, test := range tests {
		test := test
		testName := fmt.Sprintf("%d + %d = %d", test.slice1, test.slice2, test.want)

		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got := All(test.slice1, test.slice2)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
