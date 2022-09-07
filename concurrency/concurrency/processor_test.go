package concurrency_test

import (
	"testing"

	"github.com/patrickquigley102/lgwt2/concurrency/concurrency"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	t.Parallel()
	type args struct {
		data map[string]bool
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			"all true",
			args{map[string]bool{"a": true, "b": true}},
			map[string]string{"Status OK": "100%", "Status Not OK": "0%"},
		},
		{
			"some of both",
			args{map[string]bool{"a": true, "b": true, "c": false}},
			map[string]string{"Status OK": "67%", "Status Not OK": "33%"},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, concurrency.Process(test.args.data))
		})
	}
}
