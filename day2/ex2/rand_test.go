// Benchmark
// go test -bench=. example.com/day2/ex2/rand
package main

import (
	"math/big"
	"testing"
)

func BenchmarkCSRNG(b *testing.B) {
	size := b.N
	res := make([]big.Int, size)
	csrng(res[:], size)
}

func BenchmarkPRNG(b *testing.B) {
	size := b.N
	res := make([]big.Int, size)

	prng(res[:], size)
}
