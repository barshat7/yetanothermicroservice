package main


import (
	"log"
	"net"
	o "github.com/barshat7/foodieorders/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	o.RegisterOrderServiceServer(s, &server{})
	log.Printf("Starting GRPC Server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed To Serve %v", err)
	}
}