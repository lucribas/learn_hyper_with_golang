package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Timestamp    time.Time
	PreviousHash string
	Hash         string
	Data         map[string]interface{}
}

type Blockchain struct {
	GenesisBlock Block
	Chain        []Block
	Pow          int
}

func CreateBlockchain(dificuldade int) Blockchain {
	genesis := Block{
		Hash:      "000000",
		Timestamp: time.Now(),
	}
	return Blockchain{
		GenesisBlock: genesis,
		Chain:        []Block{genesis},
		Pow:          dificuldade,
	}
}

func (b Block) calcularHash() string {
	dado, _ := json.Marshal(b.Data)
	hash := sha256.Sum256(dado)
	return hex.EncodeToString(hash[:])
}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	dado := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.Chain[len(b.Chain)-1]
	bloco := Block{
		Data:         dado,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now(),
	}
	hash := bloco.calcularHash()
	fmt.Println("hash = ", hash)
	bloco.Hash = hash
	b.Chain = append(b.Chain, bloco)

}

func main() {
	fmt.Println("------------------------------------------")
	blockchain := CreateBlockchain(5)
	printJson(blockchain)

	fmt.Println("------------------------------------------")
	blockchain.addBlock("luciano", "jose", 100)
	printJson(blockchain)

}

func printJson(obj interface{}) {
	str, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	fmt.Println(string(str))

}
