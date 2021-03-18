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
		got := Hello("Chris", "")
		want := "Is the relay ready, Chris?"

		assertCorrectMessage(t, got, want)
	})

	t.Run("check relay work with freaks when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Is the relay ready, freaks?"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Romanian", func(t *testing.T) {
		got := Hello("Mihai", "Romanian")
		want := "E gata releul, Mihai?"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Mishel", "French")
		want := "Le relais est-il prêt, Mishel?"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Michael", "German")
		want := "Ist das Relais bereit, Michael?"
		assertCorrectMessage(t, got, want)
	})

}