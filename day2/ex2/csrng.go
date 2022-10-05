// CSRNG (cryptographically secure random number generator)

package main

import (
	"crypto/rand"
	"math/big"
)

func csrng(buf []big.Int, size int) {
	for i := 0; i < size; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(100))
		// fmt.Print(result)
		buf[i] = *result
	}
}
