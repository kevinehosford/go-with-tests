package main

import "testing"

func TestHello(t *testing.T) {

	// add an assertion function
	assertTestOutput := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	}

	// add a Test.Run for returning name
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Go")
		want := "Hello, Go!"

		assertTestOutput(t, got, want)
	})

	// add a Test.Run for default meesage
	t.Run("defaulting to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertTestOutput(t, got, want)
	})
}
