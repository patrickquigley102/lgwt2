package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		wantW string
	}{
		{"test", "3\n2\n1\nGo!"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			Countdown(w)
			assert.Equal(t, test.wantW, w.String())
		})
	}
}
