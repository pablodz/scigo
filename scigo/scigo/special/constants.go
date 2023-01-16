package special

// Chebyshev coefficients for exp(-x) I0(x)
// in the interval [0,8].
//
// lim(x->0){ exp(-x) I0(x) } = 1.
var A = []float64{
	-4.41534164647933937950e-18,
	3.33079451882223809783e-17,
	-2.43127984654795469359e-16,
	1.71539128555513303061e-15,
	-1.16853328779934516808e-14,
	7.67618549860493561688e-14,
	-4.85644678311192946090e-13,
	2.95505266312963983461e-12,
	-1.72682629144155570723e-11,
	9.67580903537323691224e-11,
	-5.18979560163526290666e-10,
	2.65982372468238665035e-9,
	-1.30002500998624804212e-8,
	6.04699502254191894932e-8,
	-2.67079385394061173391e-7,
	1.11738753912010371815e-6,
	-4.41673835845875056359e-6,
	1.64484480707288970893e-5,
	-5.75419501008210370398e-5,
	1.88502885095841655729e-4,
	-5.76375574538582365885e-4,
	1.63947561694133579842e-3,
	-4.32430999505057594430e-3,
	1.05464603945949983183e-2,
	-2.37374148058994688156e-2,
	4.93052842396707084878e-2,
	-9.49010970480476444210e-2,
	1.71620901522208775349e-1,
	-3.04682672343198398683e-1,
	6.76795274409476084995e-1,
}

// Chebyshev coefficients for exp(-x) sqrt(x) I0(x)
// in the inverted interval [8,infinity].
//
// lim(x->inf){ exp(-x) sqrt(x) I0(x) } = 1/sqrt(2pi).
var B = []float64{
	-7.23318048787475395456e-18,
	-4.83050448594418207126e-18,
	4.46562142029675999901e-17,
	3.46122286769746109310e-17,
	-2.82762398051658348494e-16,
	-3.42548561967721913462e-16,
	1.77256013305652638360e-15,
	3.81168066935262242075e-15,
	-9.55484669882830764870e-15,
	-4.15056934728722208663e-14,
	1.54008621752140982691e-14,
	3.85277838274214270114e-13,
	7.18012445138366623367e-13,
	-1.79417853150680611778e-12,
	-1.32158118404477131188e-11,
	-3.14991652796324136454e-11,
	1.18891471078464383424e-11,
	4.94060238822496958910e-10,
	3.39623202570838634515e-9,
	2.26666899049817806459e-8,
	2.04891858946906374183e-7,
	2.89137052083475648297e-6,
	6.88975834691682398426e-5,
	3.36911647825569408990e-3,
	8.04490411014108831608e-1,
}
