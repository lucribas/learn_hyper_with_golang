package main

import (
	"log"
	"net"

	api "example.com/day2/ex8/server/infrastructure/server/grpc"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterUserServiceServer(grpcServer, &s)

	log.Fatal(grpcServer.Serve(listener))

}
