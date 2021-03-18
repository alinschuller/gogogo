package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("checking relay work with people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Is the relay ready, Chris?"

		assertCorrectMessage(t, got, want)
	})

	t.Run("check relay work with freaks when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Is the relay ready, freaks?"

		assertCorrectMessage(t, got, want)
	})

}