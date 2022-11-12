package main

import (
	"encoding/json"
	"fmt"

	"example.com/day2/ex8/server/infrastructure/server/grpc"
	"google.golang.org/protobuf/proto"
)

type User struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
}

func ex1() {
	userGrpc := grpc.User{
		Name:        "Luciano Ribas",
		Coordinates: []float64{23.45256, 23.45266},
	}
	userGrpcBytes, _ := proto.Marshal(&userGrpc)
	fmt.Printf("%v : [% x]\n", len(userGrpcBytes), userGrpcBytes)

	userJson := User{
		Name:        "Luciano Ribas",
		Coordinates: []float64{23.45256, 23.45266},
	}
	userJsonBytes, _ := json.Marshal(&userJson)
	fmt.Printf("%v : [% x]\n", len(userJsonBytes), userJsonBytes)
}

// 33 : [0a 0d                      4c 75 63 69 61 6e 6f 20 52 69 62 61 73     12 10 4e 7a df f8 da 73 37 40 16 35 98 86 e1 73 37 40]
// 58 : [7b 22 6e 61 6d 65 22 3a 22 4c 75 63 69 61 6e 6f 20 52 69 62 61 73     22 2c 22 63 6f 6f 72 64 69 6e 61 74 65 73 22 3a 5b 32 33 2e 34 35 32 35 36 2c 32 33 2e 34 35 32 36 36 5d 7d]
