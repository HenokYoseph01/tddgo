package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Henok")
	want := "Hello, Henok!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
