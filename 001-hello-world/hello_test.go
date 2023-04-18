package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello2, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
	t.Log("Hello, world!")
}
