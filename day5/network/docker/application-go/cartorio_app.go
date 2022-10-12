// peer chaincode query -C devchannel --name cartorio -c '{"Args":["imovelGet","123213"]}'


// https://hyperledger.github.io/fabric-gateway/
// https://github.com/hyperledger/fabric-samples/blob/main/asset-transfer-basic/application-go/assetTransfer.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	// "github.com/hyperledger/fabric-sdk-go"
)

type MintImovelPayload struct {
	Inscricao string `json:"inscricao"`
	Nome      string `json:"nome"`
	Tamanho   int    `json:"tamanho"`
}

func main() {

	// contract, _ := connectContract()
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"gateway",
		"connection-org1.yaml",
	)

	time.Sleep(3 * time.Second)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	fmt.Println("BBBBBBBBBBBBBBBBB")
	time.Sleep(3 * time.Second)
	network, err := gw.GetNetwork("devchannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}
	time.Sleep(3 * time.Second)

	fmt.Println("AAAAAAAAAAAAAAAAAAA")
	contract := network.GetContract("cartorio")

	// result, err := contract.EvaluateTransaction("queryAllCars")
	// if err != nil {
	// 	fmt.Printf("Failed to evaluate transaction: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(result))

	//result, err := contract.SubmitTransaction("get", "SomeIssuer", "0001")

	if true {
		mintData := &MintImovelPayload{
			Inscricao: "343242235252-2",
			Nome:      "Apartamento 1234, Ed. Dom Pedro",
			Tamanho:   260,
		}
		b, err := json.Marshal(mintData)
		fmt.Println(string(b))
		result, err := contract.SubmitTransaction("imovel.imovelMint", string(b))
		if err != nil {
			fmt.Printf("Failed to submit transaction: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(result))

	}

	// result, err := contract.EvaluateTransaction("get", "SomeIssuer", "0001")
	// if err != nil {
	// 	fmt.Printf("Failed to submit transaction: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(result))

	// result, err := contract.EvaluateTransaction("list")
	// if err != nil {
	// 	fmt.Printf("Failed to submit transaction: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(result))

	// result, err = contract.EvaluateTransaction("queryCar", "CAR10")
	// if err != nil {
	// 	fmt.Printf("Failed to evaluate transaction: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(result))

	// _, err = contract.SubmitTransaction("changeCarOwner", "CAR10", "Archie")
	// if err != nil {
	// 	fmt.Printf("Failed to submit transaction: %s\n", err)
	// 	os.Exit(1)
	// }

	// result, err = contract.EvaluateTransaction("queryCar", "CAR10")
	// if err != nil {
	// 	fmt.Printf("Failed to evaluate transaction: %s\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(result))
}

func connectContract() (*gateway.Contract, error) {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "false")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"gateway",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("devchannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract := network.GetContract("cartorio")
	return contract, nil
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"crypto-config",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
		// "tsl",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}
