/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	helloWorldAssetContract := new(HelloWorldAssetContract)
	helloWorldAssetContract.Info.Version = "0.0.1"
	helloWorldAssetContract.Info.Description = "My Smart Contract"
	helloWorldAssetContract.Info.License = new(metadata.LicenseMetadata)
	helloWorldAssetContract.Info.License.Name = "Apache-2.0"
	helloWorldAssetContract.Info.Contact = new(metadata.ContactMetadata)
	helloWorldAssetContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(helloWorldAssetContract)
	chaincode.Info.Title = "chaincode chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from HelloWorldAssetContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
