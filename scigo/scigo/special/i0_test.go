package special

import (
	"math"
	"testing"
)

func TestI0(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		{"Test Case 1", 0, 1},
		{"Test Case 2", 1, 1.2660658777520082},
		{"Test Case 3", 2, 2.279585302336067},
		{"Test Case 4", 5, 27.239871823604442},
		{"Test Case 5", 10, 2815.716628466254},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I0(tt.x); math.Abs(got-tt.want) > 1e-14 {
				t.Errorf("I0() = %v, want %v", got, tt.want)
			}
		})
	}
}
