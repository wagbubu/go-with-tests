package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Raven")
	want := "Hello, Raven"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
