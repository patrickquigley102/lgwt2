package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testDict() Dictionary {
	return map[string]string{"test": "a test"}
}

func TestDictionary_Search(t *testing.T) {
	t.Parallel()
	type args struct {
		word string
	}
	tests := []struct {
		name      string
		d         Dictionary
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		{"word found", testDict(), args{"test"}, "a test", assert.NoError},
		{"word not found", testDict(), args{""}, "", assert.Error},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := test.d.Search(test.args.word)
			test.assertion(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestDictionary_Add(t *testing.T) {
	t.Parallel()
	type args struct {
		word string
		def  string
	}
	tests := []struct {
		name      string
		d         Dictionary
		args      args
		assertion assert.ErrorAssertionFunc
		want      string
	}{
		{"new word", testDict(), args{"a", "b"}, assert.NoError, "b"},
		{"preexisting word", testDict(), args{"test", ""}, assert.Error, "a test"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.assertion(t, test.d.Add(test.args.word, test.args.def))
			got, _ := test.d.Search(test.args.word)
			assert.Equal(t, test.want, got)
		})
	}
}
