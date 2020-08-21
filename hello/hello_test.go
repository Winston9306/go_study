package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q' ", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Winston", "")
		want := "Hello Winston"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world to people", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
}
