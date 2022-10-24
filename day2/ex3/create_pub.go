package main

// https://medium.com/coinmonks/how-to-generate-a-bitcoin-address-step-by-step-9d7fcbf1ad0b

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/ripemd160"
)

var one = big.NewInt(1)

func main() {

	//
	// (1) start with private key, or
	//
	// privateKey := "0000000000000000000000000000000000000000000000000000000000000001"
	// // privateKey := "E83385AF76B2B1997326B567461FB73DD9C27EAB9E1E86D26779F4650C5F2B75"
	// publicKey := generatePK(privateKey)

	privateKey := big.NewInt(3)
	publicKey := generatePKBig(privateKey)

	//
	// (2) start with public key
	//
	// pk := "04678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5f"
	// pk := "043cba1f4d12d1ce0bced725373769b2262c6daa97be6a0588cfec8ce1a5f0bd092f56b5492adbfc570b15644c74cc8a4874ed20dfe47e5dce2e08601d6f11f5a4"
	// publicKey := readPK(pk)

	btcWallet := EncodeWallet(publicKey)
	log.Printf("BTC Wallet = %s\n", btcWallet)

	// start := big.NewInt(9223372036854775807)
	// end := big.NewInt(18446744073709551615)

	start := big.NewInt(10)
	end := big.NewInt(1000000000000000)

	// 9223372036854775807
	// 18446744073709551615

	step := 0
	// i must be a new int so that it does not overwrite start
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {
		step++
		publicKey := generatePKBig(i)
		btcWallet := EncodeWallet(publicKey)
		if btcWallet == "16jY7qLJnxb7CHZyqBP8qca9d51gAjyXQN" {
			log.Printf("BTC privakey = %d\n", i)
			log.Printf("BTC Wallet (%d)= %s\n", i, btcWallet)
		} else {
			if step > 100000 {
				step = 0
				log.Printf("BTC privakey = %d\n", i)
				log.Printf(".")
			}
		}
	}

}

func generatePK(privateKey string) (publicKey []byte) {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	// log.Printf("  Y =%x\n", e.PublicKey.Y)
	// log.Printf("  X =%x\n", e.PublicKey.X)
	//publicKey = elliptic.Marshal(secp256k1.S256(), e.X, e.Y)
	publicKey = elliptic.MarshalCompressed(secp256k1.S256(), e.X, e.Y)
	// log.Printf("PK=%x\n", publicKey)
	// log.Printf("% x  %d\n", publicKey, len(publicKey))
	return publicKey
}

func generatePKBig(privateKey *big.Int) (publicKey []byte) {
	var e ecdsa.PrivateKey
	e.D = privateKey
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	// log.Printf("  Y =%x\n", e.PublicKey.Y)
	// log.Printf("  X =%x\n", e.PublicKey.X)
	//publicKey = elliptic.Marshal(secp256k1.S256(), e.X, e.Y)
	publicKey = elliptic.MarshalCompressed(secp256k1.S256(), e.X, e.Y)
	// log.Printf("PK=%x\n", publicKey)
	// log.Printf("% x  %d\n", publicKey, len(publicKey))
	return publicKey
}

func readPK(pk string) (publicKey []byte) {
	data, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	return data
}

func EncodeWallet(publicKey []byte) (wallet string) {
	h256 := sha256.New()
	h256.Write(publicKey)
	bs := h256.Sum(nil)
	// log.Printf("sh256 =%x\n", bs)

	hrip160 := ripemd160.New()
	hrip160.Write(bs)
	hashBytes := hrip160.Sum(nil)
	// log.Printf("rip160=%x\n", hashBytes)
	hashBytes2 := append([]byte{0x00}, hashBytes...)
	// log.Printf("rip160+=%x\n", hashBytes2)

	h256 = sha256.New()
	h256.Write(hashBytes2)
	bs = h256.Sum(nil)
	// log.Printf("sh256 =%x\n", bs)

	h256 = sha256.New()
	h256.Write(bs)
	bs = h256.Sum(nil)
	// log.Printf("sh256 =%x\n", bs)

	// log.Printf("4bytes=%x\n", bs[0:4])
	hashBytes = append(hashBytes2, bs[0:4]...)

	encoded := base58.Encode(hashBytes)
	// log.Printf("base58=%s\n", encoded)
	return encoded
}

// https://privatekeys.pw/puzzles/bitcoin-puzzle-tx
// puzzle
// https://bitcointalk.org/index.php?topic=5166284
// 16jY7qLJnxb7CHZyqBP8qca9d51gAjyXQN
// 9223372036854775807
// 18446744073709551615

// performance amd x12nucleos
// 71khashes/s
// 9223372036854775807/71000
// 129906648406405s
// 4.119.312 anos
