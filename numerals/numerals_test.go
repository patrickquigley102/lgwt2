// Package numerals is for numerals.
package numerals_test

import (
	"testing"

	"github.com/patrickquigley102/lgwt2/numerals"
	"github.com/stretchr/testify/assert"
)

func TestConvertToNumerals(t *testing.T) {
	t.Parallel()
	type args struct {
		num int
	}
	tests := []struct {
		name      string
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		{"1", args{1}, "I", assert.NoError},
		{"2", args{2}, "II", assert.NoError},
		{"4", args{4}, "IV", assert.NoError},
		{"5", args{5}, "V", assert.NoError},
		{"6", args{6}, "VI", assert.NoError},
		{"9", args{9}, "IX", assert.NoError},
		{"10", args{10}, "X", assert.NoError},
		{"11", args{11}, "XI", assert.NoError},
		{"40", args{40}, "XL", assert.NoError},
		{"50", args{50}, "L", assert.NoError},
		{"51", args{51}, "LI", assert.NoError},
		{"90", args{90}, "XC", assert.NoError},
		{"100", args{100}, "C", assert.NoError},
		{"101", args{101}, "CI", assert.NoError},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := numerals.ConvertToNumerals(test.args.num)
			test.assertion(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
