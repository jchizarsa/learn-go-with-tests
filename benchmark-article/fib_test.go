// Trying this out: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

package benchmarkarticle

import "testing"

func TestFib(t *testing.T) {
	result := Fib(6)
	expected := 8

	if result != expected {
		t.Errorf("expected %d but received %d", expected, result)
	}
}

// Skip TestX functions by passing -run=XXX, or some regex that will never match
func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
