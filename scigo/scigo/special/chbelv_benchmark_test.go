package special

import (
	"testing"
)

func BenchmarkChbevl(b *testing.B) {
	x := 3.14
	for i := 0; i < b.N; i++ {
		Chbevl(x, A, len(A))
	}
}
