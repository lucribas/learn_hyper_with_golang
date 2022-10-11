/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

func configureStub() (*MockContext, *MockStub) {
	var nilBytes []byte

	testHelloWorldAsset := new(HelloWorldAsset)
	testHelloWorldAsset.Value = "set value"
	helloWorldAssetBytes, _ := json.Marshal(testHelloWorldAsset)

	ms := new(MockStub)
	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
	ms.On("GetState", "missingkey").Return(nilBytes, nil)
	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
	ms.On("GetState", "helloWorldAssetkey").Return(helloWorldAssetBytes, nil)
	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

	mc := new(MockContext)
	mc.On("GetStub").Return(ms)

	return mc, ms
}

func TestHelloWorldAssetExists(t *testing.T) {
	var exists bool
	var err error

	ctx, _ := configureStub()
	c := new(HelloWorldAssetContract)

	exists, err = c.HelloWorldAssetExists(ctx, "statebad")
	assert.EqualError(t, err, getStateError)
	assert.False(t, exists, "should return false on error")

	exists, err = c.HelloWorldAssetExists(ctx, "missingkey")
	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
	assert.False(t, exists, "should return false when no value for key in world state")

	exists, err = c.HelloWorldAssetExists(ctx, "existingkey")
	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
	assert.True(t, exists, "should return true when value for key in world state")
}

func TestCreateHelloWorldAsset(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(HelloWorldAssetContract)

	err = c.CreateHelloWorldAsset(ctx, "statebad", "some value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.CreateHelloWorldAsset(ctx, "existingkey", "some value")
	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

	err = c.CreateHelloWorldAsset(ctx, "missingkey", "some value")
	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
}

func TestReadHelloWorldAsset(t *testing.T) {
	var helloWorldAsset *HelloWorldAsset
	var err error

	ctx, _ := configureStub()
	c := new(HelloWorldAssetContract)

	helloWorldAsset, err = c.ReadHelloWorldAsset(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
	assert.Nil(t, helloWorldAsset, "should not return HelloWorldAsset when exists errors when reading")

	helloWorldAsset, err = c.ReadHelloWorldAsset(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
	assert.Nil(t, helloWorldAsset, "should not return HelloWorldAsset when key does not exist in world state when reading")

	helloWorldAsset, err = c.ReadHelloWorldAsset(ctx, "existingkey")
	assert.EqualError(t, err, "Could not unmarshal world state data to type HelloWorldAsset", "should error when data in key is not HelloWorldAsset")
	assert.Nil(t, helloWorldAsset, "should not return HelloWorldAsset when data in key is not of type HelloWorldAsset")

	helloWorldAsset, err = c.ReadHelloWorldAsset(ctx, "helloWorldAssetkey")
	expectedHelloWorldAsset := new(HelloWorldAsset)
	expectedHelloWorldAsset.Value = "set value"
	assert.Nil(t, err, "should not return error when HelloWorldAsset exists in world state when reading")
	assert.Equal(t, expectedHelloWorldAsset, helloWorldAsset, "should return deserialized HelloWorldAsset from world state")
}

func TestUpdateHelloWorldAsset(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(HelloWorldAssetContract)

	err = c.UpdateHelloWorldAsset(ctx, "statebad", "new value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.UpdateHelloWorldAsset(ctx, "missingkey", "new value")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.UpdateHelloWorldAsset(ctx, "helloWorldAssetkey", "new value")
	expectedHelloWorldAsset := new(HelloWorldAsset)
	expectedHelloWorldAsset.Value = "new value"
	expectedHelloWorldAssetBytes, _ := json.Marshal(expectedHelloWorldAsset)
	assert.Nil(t, err, "should not return error when HelloWorldAsset exists in world state when updating")
	stub.AssertCalled(t, "PutState", "helloWorldAssetkey", expectedHelloWorldAssetBytes)
}

func TestDeleteHelloWorldAsset(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(HelloWorldAssetContract)

	err = c.DeleteHelloWorldAsset(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.DeleteHelloWorldAsset(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

	err = c.DeleteHelloWorldAsset(ctx, "helloWorldAssetkey")
	assert.Nil(t, err, "should not return error when HelloWorldAsset exists in world state when deleting")
	stub.AssertCalled(t, "DelState", "helloWorldAssetkey")
}
