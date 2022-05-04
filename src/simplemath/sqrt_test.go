package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	i := Sqrt(9)
	if i != 3 {
		t.Errorf("Sqrt(9) failed. Got %v, expected 3.", i)
	}
}
