// [_SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2)

package main

import (
	"crypto/sha256"
	"fmt"

	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

func main() {

	//s := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	// s := []byte{0x01}

	pk := "0450863AD64A87AE8A2FE83C1AF1A8403CB53F53E486D8511DAD8A04887E5B23522CD470243453A299FA9E77237716103ABC11A1DF38855ED6F2EE187E9C582BA6"
	// pk := "0000000000000000000000000000000000000000000000000000000000000001"
	data, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x  %d\n", data, len(data))

	h := sha256.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write(data)

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	fmt.Printf("%x\n", bs)

	hasher := ripemd160.New()
	hasher.Write(bs)
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashString)

}
