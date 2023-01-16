package special

import "testing"

func BenchmarkI0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		I0(5)
	}
}

func BenchmarkI0Pi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		I0(3.14)
	}
}
