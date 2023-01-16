package special

import (
	"testing"
)

func BenchmarkChbevl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Chbevl(3.14, A, len(A))
	}
}
