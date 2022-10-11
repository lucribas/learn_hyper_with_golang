/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// HelloWorldAssetContract contract for managing CRUD for HelloWorldAsset
type HelloWorldAssetContract struct {
	contractapi.Contract
}

// HelloWorldAssetExists returns true when asset with given ID exists in world state
func (c *HelloWorldAssetContract) HelloWorldAssetExists(ctx contractapi.TransactionContextInterface, helloWorldAssetID string) (bool, error) {
	data, err := ctx.GetStub().GetState(helloWorldAssetID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateHelloWorldAsset creates a new instance of HelloWorldAsset
func (c *HelloWorldAssetContract) CreateHelloWorldAsset(ctx contractapi.TransactionContextInterface, helloWorldAssetID string, value string) error {
	exists, err := c.HelloWorldAssetExists(ctx, helloWorldAssetID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", helloWorldAssetID)
	}

	helloWorldAsset := new(HelloWorldAsset)
	helloWorldAsset.Value = value

	bytes, _ := json.Marshal(helloWorldAsset)

	return ctx.GetStub().PutState(helloWorldAssetID, bytes)
}

// ReadHelloWorldAsset retrieves an instance of HelloWorldAsset from the world state
func (c *HelloWorldAssetContract) ReadHelloWorldAsset(ctx contractapi.TransactionContextInterface, helloWorldAssetID string) (*HelloWorldAsset, error) {
	exists, err := c.HelloWorldAssetExists(ctx, helloWorldAssetID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", helloWorldAssetID)
	}

	bytes, _ := ctx.GetStub().GetState(helloWorldAssetID)

	helloWorldAsset := new(HelloWorldAsset)

	err = json.Unmarshal(bytes, helloWorldAsset)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type HelloWorldAsset")
	}

	return helloWorldAsset, nil
}

// UpdateHelloWorldAsset retrieves an instance of HelloWorldAsset from the world state and updates its value
func (c *HelloWorldAssetContract) UpdateHelloWorldAsset(ctx contractapi.TransactionContextInterface, helloWorldAssetID string, newValue string) error {
	exists, err := c.HelloWorldAssetExists(ctx, helloWorldAssetID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", helloWorldAssetID)
	}

	helloWorldAsset := new(HelloWorldAsset)
	helloWorldAsset.Value = newValue

	bytes, _ := json.Marshal(helloWorldAsset)

	return ctx.GetStub().PutState(helloWorldAssetID, bytes)
}

// DeleteHelloWorldAsset deletes an instance of HelloWorldAsset from the world state
func (c *HelloWorldAssetContract) DeleteHelloWorldAsset(ctx contractapi.TransactionContextInterface, helloWorldAssetID string) error {
	exists, err := c.HelloWorldAssetExists(ctx, helloWorldAssetID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", helloWorldAssetID)
	}

	return ctx.GetStub().DelState(helloWorldAssetID)
}
