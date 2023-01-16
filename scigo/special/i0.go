package special

// Modified Bessel function of order zero
//
// SYNOPSIS:
//
// 	double x, y, i0();
//
// y = i0( x );
//
//
//
// DESCRIPTION:
//
// Returns modified Bessel function of order zero of the
// argument.
//
// The function is defined as i0(x) = j0( ix ).
//
// The range is partitioned into the two intervals [0,8] and
// (8, infinity). Chebyshev polynomial expansions are employed
// in each interval.
//
//
//
// ACCURACY:
//
// Relative error:
// arithmetic domain # trials peak rms
// IEEE 0,30 30000 5.8e-16 1.4e-16
//
//

// 						i0e.c
// Modified Bessel function of order zero,
// exponentially scaled
// SYNOPSIS:
// 		double x, y, i0e();
// y = i0e( x );
// DESCRIPTION:
// Returns exponentially scaled modified Bessel function
// of order zero of the argument.
// The function is defined as i0e(x) = exp(-|x|) j0( ix ).
// ACCURACY:
// Relative error:
// arithmetic domain # trials peak rms
// IEEE 0,30 30000 5.4e-16 1.2e-16
// See i0().
//

// Cephes Math Library Release 2.8: June, 2000
// Copyright 1984, 1987, 2000 by Stephen L. Moshier

import "math"

// I0 returns the exponentially scaled modified Bessel
// function of order zero of the argument.
func I0(x float64) float64 {
	var y float64

	if x < 0 {
		x = -x
	}
	if x <= 8.0 {
		y = (x / 2.0) - 2.0
		return (math.Exp(x) * Chbevl(y, A, 30))
	}

	return (math.Exp(x) * Chbevl(32.0/x-2.0, B, 25) / math.Sqrt(x))
}
