package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"google.golang.org/grpc"

	serializers "./lib/serializers/firestore"
	firestore "./lib/db/firestore"
	helloPb "./helloworld"
	storePb "./store"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer and store.StoreServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
// this is an example of a receiver which adds the SayHeloMethod to the server struct
func (s *server) SayHello(ctx context.Context, in *helloPb.HelloRequest) (*helloPb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloPb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// GetLineItem implements store.StoreServer
func (s *server) GetLineItem(ctx context.Context, in *storePb.LineItemRequest) (*storePb.LineItem, error) {
	log.Printf("Attempting to fetch line item: %v", in.GetId())

	client, fsctx := firestore.Connect()
	defer client.Close()

	lineItemData := firestore.FetchLineItem(client, fsctx, strconv.FormatInt(int64(in.GetId()), 10))
	log.Printf("Fetched Line item: %+v", lineItemData)

	result := serializers.SerializeLineItem(lineItemData)
	log.Printf("Serilialized Line item: %+v", result)

	return result, nil
}

func main() {
	// open up port to listen
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}

	// init grpc server
	s := grpc.NewServer()

	// register methods
	helloPb.RegisterGreeterServer(s, &server{})
	storePb.RegisterStoreServer(s, &server{})

	// start grpc server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error: Failed to serve: %v", err)
	}
}
