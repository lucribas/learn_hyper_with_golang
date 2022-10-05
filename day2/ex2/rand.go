// CSRNG (cryptographically secure random number generator)

package main

import (
	"fmt"
	"math/big"
)

func main() {
	size := 10000
	res := make([]big.Int, size)

	csrng(res[:], size)
	fmt.Println()
	for _, v := range res {
		fmt.Print(&v)
	}

	prng(res[:], size)
	fmt.Println()
	for _, v := range res {
		fmt.Print(&v)
	}
}
