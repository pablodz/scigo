package windows

import (
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestKaiserDerived(t *testing.T) {

	tests := []struct {
		name string
		M    int
		beta float64
		sym  bool
		want []float64
	}{
		{"Test Case 1: M=20, beta=0, sym=True", 20, 5.0, true, []float64{
			0.08192543559052605,
			0.1986720121208764,
			0.33959762687182604,
			0.49144233519241015,
			0.6392073430825657,
			0.7690344417185278,
			0.870910116591047,
			0.9405708117005461,
			0.9800660343057719,
			0.9966384615312127,
			0.9966384615312127,
			0.9800660343057719,
			0.9405708117005461,
			0.870910116591047,
			0.7690344417185278,
			0.6392073430825657,
			0.49144233519241015,
			0.33959762687182604,
			0.1986720121208764,
			0.08192543559052605,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KaiserBesselDerived(tt.M, tt.beta, tt.sym)
			if err != nil {
				t.Errorf("Kaiser() = %v, want %v", err, tt.want)
			}

			t.Logf("len(got) = %v", len(got))
			t.Logf("len(want) = %v", len(tt.want))
			if len(got) != len(tt.want) {
				t.Errorf("Kaiser() = %v, want %v", got, tt.want)
			}
			if !floats.EqualApprox(got, tt.want, 1e-14) {
				t.Errorf("Kaiser() = %v, want %v", got, tt.want)
			}

		})
	}
}
