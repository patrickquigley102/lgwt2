package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		{"hello", args{"PQ"}, "Hello, PQ"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			writer := &bytes.Buffer{}
			Greet(writer, test.args.name)
			assert.Equal(t, test.wantWriter, writer.String())
		})
	}
}
