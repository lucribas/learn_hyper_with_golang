// PRNG (pseudo-random number generator) with seed

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 100; i++ {
		fmt.Println(r1.Intn(99))
	}
}
