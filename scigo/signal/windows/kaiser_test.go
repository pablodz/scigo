package windows

import (
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestKaiser(t *testing.T) {
	tests := []struct {
		name string
		M    int
		beta float64
		sym  bool
		want []float64
	}{
		{
			"Test Case 1: M=25, beta=0, sym=True", 25, 0, true, []float64{
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
				1,
			},
		},
		{
			"Test Case 2: M=25, beta=12, sym=True", 25, 12, true, []float64{
				5.2773441320097665e-05,
				0.0011984540164206258,
				0.006341318682879958,
				0.021278376581138667,
				0.054759364131249516,
				0.11688691109671653,
				0.2156727447589671,
				0.35258388528729573,
				0.5188447216723607,
				0.6945160085071692,
				0.8516023747395998,
				0.9608305513095535,
				1.0,
				0.9608305513095535,
				0.8516023747395998,
				0.6945160085071692,
				0.5188447216723607,
				0.35258388528729573,
				0.2156727447589671,
				0.11688691109671653,
				0.054759364131249516,
				0.021278376581138667,
				0.006341318682879958,
				0.0011984540164206258,
				5.2773441320097665e-05,
			},
		},
		{
			"Test Case 3: M=25, beta=12, sym=False", 25, 12, false, []float64{
				5.2773441320097665e-05,
				0.0011037219720325718,
				0.005666659918028084,
				0.01870180047071839,
				0.04773662957087116,
				0.10171968770956323,
				0.18839351891936634,
				0.31069994737921397,
				0.46345418080132217,
				0.6318444677354348,
				0.7929398341014605,
				0.9203475105333124,
				0.9908477661338682,
				0.9908477661338682,
				0.9203475105333124,
				0.7929398341014605,
				0.6318444677354348,
				0.46345418080132217,
				0.31069994737921397,
				0.18839351891936634,
				0.10171968770956323,
				0.04773662957087116,
				0.01870180047071839,
				0.005666659918028084,
				0.0011037219720325718,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Kaiser(tt.M, tt.beta, tt.sym)
			if err != nil {
				t.Errorf("Kaiser() = %v, want %v", err, tt.want)
			}
			t.Logf("len(got) = %v", len(got))
			t.Logf("len(want) = %v", len(tt.want))
			if !floats.EqualApprox(got, tt.want, 1e-14) {
				t.Errorf("Kaiser() = %v, want %v", got, tt.want)
			}
		})
	}
}
