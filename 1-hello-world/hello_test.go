package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("PQ")
	want := "Hello PQ"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
