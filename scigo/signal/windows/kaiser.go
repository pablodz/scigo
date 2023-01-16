package windows

import (
	"fmt"
	"math"

	"github.com/pablodz/scigo/scigo/special"
	"gonum.org/v1/gonum/floats"
)

const (
	KaiserBesselNotSymmetricError = "scigo: Kaiser-Bessel window only implemented for symmetric windows"
	NumPointsEvenError            = "scigo: Kaiser-Bessel Derived windows are only defined for even number "
)

// Kaiser returns the Kaiser window
// Parameters:
//   - M: the number of points in the window
//   - beta: the beta parameter, used if sym is false.
//   - sym: if true, generates a symmetric window, for use in filter design.
//     Otherwise, generates a periodic window, for use in spectral analysis.
//
// Returns:
//   - out: the window, with the maximum value normalized to 1 (though the value 1
//     does not appear if sym is True)
func Kaiser(M int, beta float64, sym bool) ([]float64, error) {
	ok, err := _len_guards(M)
	if err != nil {
		return nil, err
	}

	if ok {
		// generate ones of M size
		ones := make([]float64, M)
		for i := range ones {
			ones[i] = 1
		}
		return ones, nil
	}

	M, needsTrunc := _extend(M, sym)

	arange := make([]float64, M)
	n := floats.Span(arange, 0, float64(M-1))
	alpha := float64(M-1) / 2.0

	w := []float64{}

	for _, v := range n {
		num := special.I0(beta * math.Sqrt(1-math.Pow((v-alpha)/alpha, 2.0)))
		denom := special.I0(beta)
		w = append(w, num/denom)
	}

	w = _truncate(w, needsTrunc)

	return w, nil
}

// KaiserBesselDerived returns the Kaiser-Bessel window
// Parameters:
//   - M: the number of points in the window
//   - beta: the beta parameter, used if sym is false.
//   - sym: if true, generates a symmetric window, for use in filter design.
//     Otherwise, generates a periodic window, for use in spectral analysis.
//
// Returns:
//   - out: the window, with the maximum value normalized to 1 (though the value 1
//     does not appear if sym is True)
func KaiserBesselDerived(M int, beta float64, sym bool) ([]float64, error) {
	if !sym {
		return nil, fmt.Errorf(KaiserBesselNotSymmetricError)
	} else if M < 1 {
		return []float64{}, nil
	} else if M%2 == 0 {
		return nil, fmt.Errorf(NumPointsEvenError)
	}

	kaiserWindow := make([]float64, M/2+1)
	kaiserWindow = kaiserWindow[:M/2+1]

	return kaiserWindow, nil
}
