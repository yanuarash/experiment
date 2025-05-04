package main

import "testing"

func TestHello(t *testing.T) {
	hello := Hello()
	expected := "Hello, world"

	if hello != expected {
		t.Errorf("got %q want %q", hello, expected)
	}
}
