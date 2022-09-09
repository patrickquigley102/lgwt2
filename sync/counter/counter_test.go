// Package counter is used for counting.
package counter_test

import (
	"testing"

	"github.com/patrickquigley102/lgwt2/sync/counter"
	"github.com/stretchr/testify/assert"
)

func TestCounter_Value(t *testing.T) {
	t.Parallel()
	type fields struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"success", fields{1}, 1},
		{"success", fields{2}, 2},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := counter.NewCounter(test.fields.value)
			assert.Equal(t, test.want, c.Value())
		})
	}
}

func TestCounter_Inc(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		init      int
		noOfCalls int
		want      int
	}{
		{"start at 0, inc 2", 0, 2, 2},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := counter.NewCounter(test.init)
			for test.noOfCalls > 0 {
				c.Inc()
				test.noOfCalls--
			}
			assert.Equal(t, test.want, c.Value())
		})
	}
}
