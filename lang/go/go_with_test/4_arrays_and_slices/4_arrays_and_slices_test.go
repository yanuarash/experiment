package main

import "testing"

func TestSumt(t *testing.T) {
	numbers := [3]int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
