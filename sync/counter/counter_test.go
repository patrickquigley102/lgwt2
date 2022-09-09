// Package counter is used for counting.
package counter_test

import (
	"sync"
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
		{"start at 5, inc 20", 5, 20, 25},
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

func TestCounter_Inc_Concurrent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		init      int
		noOfCalls int
		want      int
	}{
		{"start at 0, inc 20", 0, 20, 20},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			counter := counter.NewCounter(test.init)
			var wait sync.WaitGroup
			wait.Add(test.noOfCalls)

			for test.noOfCalls > 0 {
				test.noOfCalls--
				go func() {
					counter.Inc()
					wait.Done()
				}()
			}
			wait.Wait()

			assert.Equal(t, test.want, counter.Value())
		})
	}
}
