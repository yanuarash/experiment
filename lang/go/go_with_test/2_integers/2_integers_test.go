package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	expected := 3
	if sum != expected {
		t.Errorf("expected '%d' but got '%d", expected, sum)
	}
}

func TestAddWithMessage(t *testing.T) {
	t.Run("1 plus 3 should return 4", func(t *testing.T) {
		sum := Add(1, 3)
		expected := 4
		if sum != expected {
			t.Errorf("expected '%d' but got '%d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
