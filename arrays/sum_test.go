// Package sum for learnings
package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3 Factorial", args{[]int{1, 2, 3}}, 6},
		{"Empty slice", args{[]int{}}, 0},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, Sum(test.args.slice))
		})
	}
}

func TestAll(t *testing.T) {
	t.Parallel()
	type args struct {
		slices [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty slices", args{[][]int{{}, {}}}, []int{0, 0}},
		{"three slices", args{[][]int{{1, 2}, {3, 4}, {5, 6}}}, []int{3, 7, 11}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, All(test.args.slices...))
		})
	}
}
