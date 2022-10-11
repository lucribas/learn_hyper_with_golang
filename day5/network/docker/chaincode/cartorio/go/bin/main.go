package main

import (
	"cartorio"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	err := shim.Start(cartorio.NewCC())
	if err != nil {
		fmt.Printf("Error starting Cartorio chaincode: %s", err)
	}
}
