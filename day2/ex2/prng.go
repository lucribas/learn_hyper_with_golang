// CSRNG (cryptographically secure random number generator)

package main

import (
	"math/big"
	"math/rand"
)

func prng(buf []big.Int, size int) {
	for i := 0; i < size; i++ {
		result := big.NewInt((rand.Int63()))
		buf[i] = *result
	}
}
