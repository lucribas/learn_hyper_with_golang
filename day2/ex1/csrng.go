// CSRNG (cryptographically secure random number generator)

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Generate 20 random numbers with exclusive max of 100.
	// ... So max value returned is 99.
	//     All values returned are 0 or greater as well.
	for i := 0; i < 100; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(100))
		fmt.Println(result)
	}
}
