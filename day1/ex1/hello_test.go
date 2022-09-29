package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "HELLO"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
