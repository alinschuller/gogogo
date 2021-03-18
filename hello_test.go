package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Is the relay ready? We're getting there"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
