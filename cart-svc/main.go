package main

import (
	s "github.com/barshat7/foodiecart/cartgrpc"
	"net"
	"log"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	_server := grpc.NewServer()
	s.RegisterCartServiceServer(_server, &s.Server{})
	log.Printf("Starting GRPC Server")
	if err := _server.Serve(lis); err != nil {
		log.Fatalf("Failed To Serve %v", err)
	}
}