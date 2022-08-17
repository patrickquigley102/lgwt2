package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("with name", func(t *testing.T) {
		got := Hello("PQ", "english")
		want := "Hello PQ"

		assertCorrectString(t, got, want)
	})
	t.Run("without name", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"

		assertCorrectString(t, got, want)
	})
	t.Run("Spanish, without name", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola World"

		assertCorrectString(t, got, want)
	})
	t.Run("French, without name", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bounjour World"

		assertCorrectString(t, got, want)
	})
}

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
