// [_SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2)

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Curso de blockchain!"

	h := sha256.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
