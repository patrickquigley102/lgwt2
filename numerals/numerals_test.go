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
