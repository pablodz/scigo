package special

import (
	"math"
	"testing"
)

func TestChbevl(t *testing.T) {
	x := 2.0
	array := []float64{
		-4.41534164647933937950e-18,
		3.33079451882223809783e-17,
	}
	n := 2
	expected := 1.2238630947631852e-17
	result := Chbevl(x, array, n)
	if math.Abs(expected-result) > 1e-15 {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
