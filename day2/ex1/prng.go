// Example PRNG (pseudo-random number generator)

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(99))
	}
}
