package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("b")
	expected := "bbbbb"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)

	}
}
