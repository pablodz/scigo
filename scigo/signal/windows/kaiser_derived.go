package windows

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/floats"
)

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
	} else if M%2 != 0 {
		return nil, fmt.Errorf(NumPointsEvenError)
	}

	M2 := int(math.Ceil(float64(M)/2.0)) + 1
	kaiserWindow, err := Kaiser(M2, beta, sym)
	if err != nil {
		return nil, err
	}

	kaiserWindowCum := floats.CumSum(kaiserWindow, kaiserWindow)
	lattKaiserWindowCumValue := kaiserWindowCum[len(kaiserWindowCum)-1]
	kaiserWindowCum = kaiserWindowCum[:len(kaiserWindowCum)-1]

	for i := range kaiserWindowCum {
		num := kaiserWindowCum[i]
		den := lattKaiserWindowCumValue
		kaiserWindowCum[i] = math.Sqrt(num / den)
	}

	fmt.Println(kaiserWindowCum)

	w := make([]float64, len(kaiserWindowCum)*2)
	copy(w, kaiserWindowCum)
	floats.Reverse(kaiserWindowCum)
	copy(w[len(kaiserWindowCum):], kaiserWindowCum)

	return w, nil
}
